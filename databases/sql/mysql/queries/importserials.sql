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
