DROP TABLE IF EXISTS `import_pallets_serials`;

CREATE TABLE `import_pallets_serials` (
  pallet varchar(20),
  masterbox varchar(20),
  serial_number varchar(20),
  part_number varchar(20),
  uuid varchar(20),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE DQC41CS
(
  CARTON_NO varchar(10) NOT NULL,
  MODEL varchar(7),
  LINE_NO varchar(6),
  QTY int(2),
  INSERT_DT DATE,
  INSERT_UID varchar(8),
  LAST_UPD DATE,
  UID1 varchar(8),
  COLOR varchar(20),
  CUSTOMER varchar(20),
  REF_JO_NO varchar(20),
  JO_NO varchar(9),
  SW_REV varchar(20),
  PACKING_DATE DATE,
  UID_NO varchar(10),
  CUST_PO_NO varchar(20),
  ECO_NO varchar(20),
  CUST_CARTON_NO varchar(50),
  ACT_QTY int(12),
  STATION_ID varchar(10),
  WEIGHT varchar(20),
  IP varchar(20),
  STATE int(38),
  MO_NO varchar(9),
  VIA varchar(5)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE DQC41CS1
(
CARTON_NO varchar(10) NOT NULL,
SER_NO varchar(32) NOT NULL,
INSERT_UID varchar(8),
UID1 varchar(8)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
