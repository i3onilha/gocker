-- name: GetOracleDataSource :one
SELECT
  src.*
FROM
  oracle_customers cust
  INNER JOIN oracle_datasources src ON cust.dbname = src.dbname
WHERE cust.customer = ?;

-- name: Create :execresult
INSERT IGNORE INTO labels () VALUES ();

-- name: GetByID :one
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.id = ?
ORDER BY labels_data.created_at DESC
LIMIT 1;

-- name: ListByModel :many
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.customer = ?
  AND labels_data.model = ?
  AND labels_data.part_number = ''
  AND labels_data.created_at IN(
    SELECT MAX(labels_data.created_at)
    FROM labels_data
    GROUP BY labels_data.id
  )
ORDER BY labels_data.created_at DESC;

-- name: ListByParts :many
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.part_number = ?
  AND labels_data.created_at IN(
    SELECT MAX(labels_data.created_at)
    FROM labels_data
    GROUP BY labels_data.id
  )
ORDER BY labels_data.created_at DESC;

-- name: ListByModelAndStationAndDpi :many
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.customer = ?
  AND labels_data.model = ?
  AND labels_data.station = ?
  AND labels_data.dpi = ?
  AND labels_data.part_number = ''
  AND labels_data.created_at IN(
    SELECT MAX(labels_data.created_at)
    FROM labels_data
    GROUP BY labels_data.id
  )
ORDER BY labels_data.created_at DESC;

-- name: ListByPartsAndStationAndDpi :many
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.customer = ?
  AND labels_data.part_number = ?
  AND labels_data.station = ?
  AND labels_data.dpi = ?
  AND labels_data.created_at IN(
    SELECT MAX(labels_data.created_at)
    FROM labels_data
    GROUP BY labels_data.id
  )
ORDER BY labels_data.created_at DESC;

-- name: ListPaginate :many
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.created_at IN(
    SELECT MAX(labels_data.created_at)
    FROM labels_data
    GROUP BY labels_data.id
  )
ORDER BY labels_data.created_at DESC
LIMIT ? OFFSET ?;

-- name: Update :execresult
INSERT INTO labels_data (id, name, customer, model, part_number, station, dpi, label, setup, sql_queries, author)
  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteByID :exec
INSERT INTO labels_deletes (id) VALUES (?);
