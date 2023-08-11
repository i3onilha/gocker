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
INSERT INTO labels_data (id, customer, family, model, part_number, order_number, line, station, dpi, label, author)
  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteByID :exec
INSERT INTO labels_deletes (id) VALUES (?);
