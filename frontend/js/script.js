function GET(name){
    let DataGet = {};
    //poluchaem stroky get parametrov
    let strGET = window.location.search.replace('?', '');
   
    let parts = strGET.split('&');
    
    for(i in parts){
        //razbili na paru kluch zhachenie
        let items = parts[i].split('=');
        DataGet[items[0]] = items[1];
        
    }
    return DataGet[name];
}

function replaceData(template,  data, fields){
    
    //в цикле заменяем все заглушки в шабооне
    for(let i = 0; i < fields.length; i++){
        
        template = replaceAll(template,'${'+ fields[i] + '}', data[fields[i]])
    }
    return template;
}

function escapeRegExp(string) {
    return string.replace(/[.*+?^${}()|[\]\\]/g, '\\$&'); // $& means the whole matched string
}

function replaceAll(str, find, replace) {
    return str.replace(new RegExp(escapeRegExp(find), 'g'), replace);
}
function countOnLoad(){
    let cart = localStorage.getItem('cart')
    console.log("polchili" + cart);
    if(cart == 'undefined' || cart == null){
        //esli pusta to sozdaem ee v vide pustogo massiva
        cart = [];
        console.log('tam nichego ne bylo')
    }
    else{
        cart = JSON.parse(cart);
    }
    let count = cart.length
    document.getElementById('cart-count').innerHTML = count;
}

function requestToData(url){
    let xhr = new XMLHttpRequest();
    xhr.open('GET', url, false);
    xhr.send();
    return xhr.responseText;
}
function insertCountGoods(){
    let url= "http://localhost:8090/app/getCount";
    let xhr = new XMLHttpRequest();
    xhr.open("GET", url, false);
    xhr.send();
    let goodsCount = JSON.parse(xhr.responseText);
    for(i in goodsCount){
        let y = parseInt(i)+ 1;
        let id = "category_" + y;
        console.log(goodsCount[i])
        document.getElementById(id).innerHTML =  goodsCount[i].count;
    }
}
function issetToken(){
    console.log(document.cookie);
    let token = document.cookie.replace(/(?:(?:^|.*;\s*)user_id\s*\=\s*([^;]*).*$)|^.*$/, "$1");
    if(token != ""){
        let data = {
            Token: token
        }
        let xhr = new XMLHttpRequest();
        xhr.open("POST", "http://localhost:8090/app/check_tn", false)
        xhr.send(JSON.stringify(data))
        let response = JSON.parse(xhr.responseText);
        console.log(response)
        console.log(token)
        if(response.status == true){
            document.querySelector(".log_basket-link").innerHTML = "Привет, " + `<div class="dropdown">
                                                                                    <div class="dropdown-btn"><b>${response.name}</b></div>
                                                                                    <div class="dropdown-content">
                                                                                        <div><a href="user_acc.html">Профиль</a></div>
                                                                                        <div><a onclick="logOutAcc()"><b>Выйти</b></a></div>
                                                                                    </div>
                                                                                </div>`;
            document.cookie = "token_correct=" + response.status;
            document.cookie = "user_name=" + response.name;                                                    
        }
        else{
            document.querySelector(".log_basket-link").innerHTML = `<a href="signIn.html" class="log-link">Войти</a></span>`
            alert("fuck you")
            document.cookie = "token_correct=" + "";
            document.cookie = "user_id=" + "";
            document.cookie = "user_name=" + "";  
        }
    }
    else{
            document.querySelector(".log_basket-link").innerHTML = `<a href="signIn.html" class="log-link">Войти</a></span>`
            document.cookie = "token_correct=" + "";
            document.cookie = "user_id=" + "";
            document.cookie = "user_name=" + "";  
        }
}

function logOutAcc(){
    document.cookie = "user_id=" + "";
    document.cookie = "user_name=" + "";
    document.cookie = "token_correct=" + "";   
    window.location.reload();
}