package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"crypto/md5"

	_ "github.com/go-sql-driver/mysql"
)

type Good struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Photo       string `json:"photo"`
	Articul     string `json:"articul"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Sizes       string `json:"sizes"`
	CategoryId  int		
	Isnew 		int
	Count_likes	int
}
type CategoryCount struct{
	Count int `json:"count"`
}

type UserReg struct {
	Login string 
	Email string 
	Pass  string
}
type UserLog struct {
	Login string 
	Pass  string
}
type UserTrue struct {
	Status bool		`json:"status"`
	Name   string	`json:"name"`
}
type Token struct {
	Token string
}
type ResponseStat struct {
	Status bool `json:"status"`
	Token  string `json:"token"`
}
type Order struct {
	Token 	string
	Order	string
}
func main() {

	http.HandleFunc("/app/get", func(w http.ResponseWriter, r *http.Request) {
		//разрешить кросдоменные запросы
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		//создаем подключение к базе данных
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic")

		//получаем get параметр ids
		categoryFilter := ""
		categoryId := r.URL.Query().Get("category_id")
		if categoryId != "" {
			categoryFilter = " AND `category_id` = " + categoryId
		}
		newFilter := ""
		is_new := r.URL.Query().Get("is_new")
		if is_new == "1" {
			newFilter = " AND `is_new` = 1" 
		}
		bestFilter := ""
		is_best := r.URL.Query().Get("is_best")
		if is_best == "1" {
			bestFilter = " AND `count_likes` > 10" 
		}
		idFilter := ""
		id := r.URL.Query().Get("id")
		if id != "" {
			idFilter = " AND `id` IN(" + id + ") "
		}

		//отправляем запрос в  БД
		result, _ := db.Query("SELECT * FROM inordic.goods WHERE 1" + categoryFilter + idFilter + newFilter + bestFilter)

		//получаем и выводим результат
		goods := []Good{}

		for result.Next() {

			//создаем пустой объект товара
			good := Good{}

			//переливаем в него данные из строки
			result.Scan(&good.Id, &good.Title, &good.Photo, &good.Articul, &good.Price, &good.Description, &good.Sizes, &good.CategoryId, &good.Isnew, &good.Count_likes)

			//добавляем наполненный товар в коробку с товарами
			goods = append(goods, good)
		}

		//кодируем в json
		jsonData, _ := json.Marshal(goods)

		//выводим на экран
		fmt.Fprintf(w, string(jsonData))
	})
	http.HandleFunc("/app/getCount", func(w http.ResponseWriter, r *http.Request) {
		//разрешаем посещать наш сайт всем 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//подключаемся к нашей базе данных 
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic") 
		//создаем массив, куда мы положим наше количество товаров
		count := []CategoryCount{}
		//создаем ассоциативный массив ключ(номер) - поля в бд
		list := map[int]string {
			1 : "category_id",
			2 : "category_id",
			3 : "category_id",
			4 : "is_new",
			5 : "count_likes",
		}
		//в цикле получаем количество товаров каждой из категорий и кладем их в массив
		for i := 0; i < len(list); i++ {
			y := strconv.Itoa(i + 1)
			sign := " = "
			if list[i + 1] == "is_new" {
				y = strconv.Itoa(1)
			}
			if list[i + 1] == "count_likes" {
				sign = " > "
				y = strconv.Itoa(10)
			}
			result, _ := db.Query("SELECT COUNT(*) FROM `goods` WHERE `" + list[i + 1] + "`" + sign + y)
			for result.Next() {
				//создаем пустую структуру
				countSre := CategoryCount{}
				//кладем в нее данные
				result.Scan(&countSre.Count)
				//кладем структуру в массив
				count = append(count, countSre)
			}
		}	
		//переводим массив в джейсон
		countJson, _ := json.Marshal(count)
		//выводим на экран результат переведя его в строку
		fmt.Fprintf(w, string(countJson))
	})
	http.HandleFunc("/app/registration", func(w http.ResponseWriter, r *http.Request) {
		//разрешаем посещать наш сайт всем 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//подключаемся к нашей базе данных 
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic") 
		//считать данные из тела запроса
		body, _ := ioutil.ReadAll(r.Body)
		//перелисть их в структуру
		regData := UserReg{} 
		json.Unmarshal(body, &regData)
		//проверяем пользователя на оригинальность
		countCommon := 0
		resultCount, _ := db.Query("SELECT COUNT(*) FROM `users` WHERE `login` = ? OR `email` = ?", regData.Login, regData.Email)
		for resultCount.Next(){
			resultCount.Scan(&countCommon)
		}
		respStat := ResponseStat{}
		if countCommon == 0{
			userToken := md5.Sum([]byte(regData.Login))
			userPass := md5.Sum([]byte(regData.Pass))
			userPassLine := fmt.Sprintf("%x", userPass)
			userTokenLine := fmt.Sprintf("%x", userToken)
			_, err := db.Query("INSERT INTO `users`(`login`, `email`, `password`, `token`) VALUES(?,?,?,?)", regData.Login, regData.Email, userPassLine, userTokenLine)
			if err != nil{
				fmt.Println(err)
			}
			//вернуть результат
			respStat = ResponseStat{Status: true, Token: userTokenLine}
		} else {
			respStat = ResponseStat{Status: false, Token: "0"}
		}
		jsonrespStat, _ := json.Marshal(respStat)
		fmt.Fprintf(w, string(jsonrespStat))

	})
	http.HandleFunc("/app/signin", func(w http.ResponseWriter, r *http.Request) {
		//разрешаем посещать наш сайт всем 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//подключаемся к нашей базе данных 
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic") 
		//считать данные из тела запроса
		body, _ := ioutil.ReadAll(r.Body)
		//перелисть их в структуру
		logData := UserLog{} 
		json.Unmarshal(body, &logData)
		//проверяем пользователя на оригинальность
		countCommon := 0
		userPass := md5.Sum([]byte(logData.Pass))
		userPassLine := fmt.Sprintf("%x", userPass)
		resultCount, _ := db.Query("SELECT COUNT(*) FROM `users` WHERE (`login` = ? OR `email` = ?) AND `password` = ?", logData.Login, logData.Login, userPassLine)
		for resultCount.Next(){ 
			resultCount.Scan(&countCommon)
		}
		respStat := ResponseStat{}
		if countCommon == 1{
			token := ""
			tokenData, _ := db.Query("SELECT token FROM `users` WHERE `login` = ? OR `email` = ?", logData.Login, logData.Login)
			for tokenData.Next(){
				tokenData.Scan(&token)
			}
			respStat = ResponseStat{Status: true, Token: token}
			//вернуть результат
		} else {
			respStat = ResponseStat{Status: false, Token: "0"}
		}
		jsonrespStat, _ := json.Marshal(respStat)
		fmt.Fprintf(w, string(jsonrespStat))

	})
	http.HandleFunc("/app/check_tn", func(w http.ResponseWriter, r *http.Request) {
		//разрешаем посещать наш сайт всем 
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//подключаемся к нашей базе данных 
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic") 
		//считать данные из тела запроса
		body, _ := ioutil.ReadAll(r.Body)
		//перелисть их в структуру
		token := Token{}
		json.Unmarshal(body, &token)
		//проверяем пользователя на оригинальность
		result, _ := db.Query("SELECT COUNT(*) FROM `users` WHERE `token`= ? ", token.Token)
		resultData := 0
		for result.Next(){
			result.Scan(&resultData)
		}
		response := UserTrue{}
		if resultData != 0{
			resultName, _ := db.Query("SELECT `login` FROM `users` WHERE `token` = ?", token.Token)
			name := ""
			for resultName.Next(){
				resultName.Scan(&name)
			}
			response = UserTrue{Status: true, Name: name}
			responseJson, _ := json.Marshal(response)
			fmt.Fprintf(w, string(responseJson))
		}	else{
			response = UserTrue{Status: false, Name: ""}
			responseJson, _ := json.Marshal(response)
			fmt.Fprintf(w, string(responseJson))
		}

	})
	http.HandleFunc("/app/order", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		db, _ := sql.Open("mysql", "root:nordic123@tcp(database:3306)/inordic") 
		//считать данные из тела запроса
		body, _ := ioutil.ReadAll(r.Body)
		order := Order{}
		json.Unmarshal(body, order)
		orderJson := ""
		ordersArray := []
		result, _ := db.Query("SELECT `orders` FROM `users` WHERE `token`=?", order.Token)
		for result.Next({
			result.Scan(&ordersJson)
		}
		json.Unmarshal(orderJson, ordersArray)
		ordersArray = append(ordersArray, order.Order)
		jsonOrders, _ := json.Marshal(ordersArray)
		db.Exec("INSERT INTO `users`(`orders`) VALUES(?) WHERE `token`=?", string(jsonOrders), order.Token) 

	})
	http.ListenAndServe(":8000", nil)
}
