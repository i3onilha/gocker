DROP TABLE IF EXISTS `import_pallets_serials`;

CREATE TABLE `import_pallets_serials` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  pallet varchar(20),
  masterbox varchar(20),
  serial_number varchar(20),
  part_number varchar(20),
  uuid varchar(20),
  `created_at` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
