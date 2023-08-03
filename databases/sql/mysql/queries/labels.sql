-- name: CreateLabel :execresult
INSERT IGNORE INTO labels () VALUES ();

-- name: GetLabel :one
SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.id = ?
ORDER BY labels_data.created_at DESC
LIMIT 1;

-- name: GetLabelList :many
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

-- name: UpdateLabel :execresult
INSERT INTO labels_data (id, customer, family, model, part_number, station, label)
  VALUES(?, ?, ?, ?, ?, ?, ?);

-- name: DeleteLabel :exec
INSERT INTO labels_deletes (id) VALUES (?);
