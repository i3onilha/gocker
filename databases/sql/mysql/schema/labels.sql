DROP DATABASE `dbdev`;

CREATE DATABASE `dbdev`;

USE `dbdev`;

DROP TABLE IF EXISTS `labels`;

CREATE TABLE `labels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

LOCK TABLES `labels` WRITE;
INSERT INTO `labels` (id) VALUES (1),(2),(3),(4);
UNLOCK TABLES;

DROP TABLE IF EXISTS `labels_data`;

CREATE TABLE `labels_data` (
  `id` int(11) NOT NULL,
  `customer` varchar(32) NOT NULL,
  `family` varchar(8) NOT NULL,
  `model` varchar(16) NOT NULL,
  `part_number` varchar(16) NOT NULL,
  `order_number` varchar(9) NOT NULL,
  `line` varchar(8) NOT NULL,
  `station` varchar(32) NOT NULL,
  `label` text NOT NULL,
  `author` varchar(8) NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`, `created_at`),
  CONSTRAINT `labels_data_ibfk_1` FOREIGN KEY (`id`) REFERENCES `labels` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `labels_data` WRITE;
INSERT INTO `labels_data` VALUES (1,'Customer','Family','Model','PartNumber', '17B006501', '01' , 'Station','Label', 'bc0g7190','2023-08-04 23:57:16'),(2,'Customer','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7191','2023-08-04 23:57:58'),(3,'Customer','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7192','2023-08-04 23:57:59'),(4,'Customer','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7193','2023-08-04 23:59:49'),(3,'Modified','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7194','2023-08-06 23:59:49'),(1,'Modified','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7195','2023-09-06 23:59:49'),(4,'Modified 2','Family','Model','PartNumber', '17B006501', '01', 'Station','Label', 'bc0g7196','2023-09-06 23:59:49');
UNLOCK TABLES;

DROP TABLE IF EXISTS `labels_deletes`;

CREATE TABLE `labels_deletes` (
  `id` int(11) NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  KEY `id` (`id`),
  CONSTRAINT `labels_deletes_ibfk_1` FOREIGN KEY (`id`) REFERENCES `labels_data` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `labels_deletes` WRITE;
INSERT INTO `labels_deletes` VALUES (1,'2023-08-04 23:57:16'),(2,'2023-08-04 23:57:58');
UNLOCK TABLES;
