-- name: ListSerialsByPallet :many
SELECT
  a.*
FROM
  import_pallets_serials a
WHERE a.pallet = ?
ORDER BY a.pallet ASC, a.created_at DESC;
