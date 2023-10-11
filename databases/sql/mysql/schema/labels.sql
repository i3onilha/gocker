DROP DATABASE `dbdev`;

CREATE DATABASE `dbdev`;

USE `dbdev`;

DROP TABLE IF EXISTS `labels`;

CREATE TABLE `labels` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `labels_data`;

CREATE TABLE `labels_data` (
  `id` int(11) NOT NULL,
  `name` varchar(32) NOT NULL,
  `customer` varchar(32) NOT NULL,
  `model` varchar(16) NOT NULL,
  `part_number` varchar(16) NOT NULL,
  `station` varchar(32) NOT NULL,
  `dpi` int(3) NOT NULL,
  `label` text NOT NULL,
  `setup` text NOT NULL,
  `sql_queries` text NOT NULL,
  `author` varchar(8) NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`, `created_at`),
  CONSTRAINT `labels_data_ibfk_1` FOREIGN KEY (`id`) REFERENCES `labels` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS `labels_deletes`;

CREATE TABLE `labels_deletes` (
  `id` int(11) NOT NULL,
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  KEY `id` (`id`),
  CONSTRAINT `labels_deletes_ibfk_1` FOREIGN KEY (`id`) REFERENCES `labels_data` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
