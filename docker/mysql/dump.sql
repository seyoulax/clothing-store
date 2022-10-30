-- Adminer 4.8.1 MySQL 8.0.30 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `photo` varchar(1000) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `articul` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `price` int DEFAULT NULL,
  `description` text NOT NULL,
  `sizes` varchar(150) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL,
  `category_id` tinyint DEFAULT NULL,
  `is_new` tinyint DEFAULT NULL,
  `count_likes` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `goods` (`id`, `title`, `photo`, `articul`, `price`, `description`, `sizes`, `category_id`, `is_new`, `count_likes`) VALUES
(1,	'Куртка синяя',	'img/catalog/1.jpg',	'000001',	5400,	'Крутая вещь',	'[\"S\", \"M\", \"L\"]',	1,	1,	6),
(2,	'Кожанная черная',	'img/catalog/4.jpg',	'000002',	22500,	'Крутая вещь',	'[\"S\", \"M\", \"L\"]',	2,	1,	45),
(3,	'Куртка с карманами',	'img/catalog/3.png',	'000003',	9200,	'Крутая вещь',	'[\"S\", \"M\", \"L\"]',	3,	0,	5),
(4,	'Джинсы голубые',	'img/catalog/12.jpg',	'000004',	5600,	'Классная вещь',	'[\"S\", \"M\", \"L\"]',	1,	1,	10055);

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `goods` varchar(10000) NOT NULL,
  `adress` varchar(100) NOT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `orders` (`order_id`, `goods`, `adress`, `user_id`) VALUES
(1,	'[{\"itemid\":\"4\",\"size\":\"S\"}]',	'г.Москва',	2),
(2,	'[{\"itemid\":\"4\",\"size\":\"S\"}]',	'',	2);

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `login` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(100) NOT NULL,
  `token` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

INSERT INTO `users` (`id`, `login`, `email`, `password`, `token`) VALUES
(1,	'danis',	'danis@gamil.com',	'81dc9bdb52d04dc20036dbd8313ed055',	'257ffa69697c2d5144f0b4b76b51ae95'),
(2,	'danis1',	'dfsfsd',	'827ccb0eea8a706c4c34a16891f84e7b',	'd7a032f32a2bfc23f6f901a7916fa2a7'),
(3,	'danis122',	'sddaa',	'8af6cbc147991af012674acf00089a62',	'96375e2e82a6ceb3ba9493aacaa61fe2'),
(4,	'dssf',	'sdssdsa',	'824adc61faf08c32ef279ee37320c6ab',	'07ec6f4b02db199ec1ea743144ebbd04'),
(5,	'',	'',	'd41d8cd98f00b204e9800998ecf8427e',	'd41d8cd98f00b204e9800998ecf8427e');

-- 2022-09-15 18:59:04