DROP TABLE IF EXISTS `import_pallets_serials`;

CREATE TABLE `import_pallets_serials` (
  pallet varchar(20),
  masterbox varchar(20),
  serial_number varchar(20),
  part_number varchar(20),
  uuid varchar(20),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
