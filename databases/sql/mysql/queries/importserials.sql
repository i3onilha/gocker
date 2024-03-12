-- name: GetByPallet :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.pallet = ?
ORDER BY a.pallet ASC, a.created_at DESC;

-- name: GetByMasterBox :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.masterbox = ?
ORDER BY a.pallet ASC, a.created_at DESC;

-- name: GetBySerialNumber :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.serial_number = ?
ORDER BY a.pallet ASC, a.created_at DESC;

-- name: GetByPartNumber :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.part_number = ?
ORDER BY a.pallet ASC, a.created_at DESC;

-- name: GetByUUID :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.part_number = ?
ORDER BY a.pallet ASC, a.created_at DESC;

-- name: Create :execresult
INSERT INTO import_pallets_serials (pallet, masterbox, serial_number, part_number, uuid)
  VALUES(?, ?, ?, ?, ?);

-- name: GetPalletAlreadyDone :many
SELECT
  c1.carton_no,
  c1.ser_no,
  t.*
FROM
  import_pallets_serials t
  LEFT JOIN dqc41cs                 c ON t.masterbox = c.cust_carton_no
  LEFT JOIN dqc41cs1                c1 ON t.serial_number = c1.ser_no
  AND c.carton_no = c1.carton_no
WHERE
  t.pallet = ?
  AND ( c1.carton_no IS NOT NULL
  OR c1.ser_no IS NOT NULL );
