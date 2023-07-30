-- name: CreateLabel :exec
INSERT IGNORE INTO labels (uuid) VALUES (uuid());

-- name: GetLabel :one
SELECT
  labels_data.*
FROM
  labels_data
  INNER JOIN labels ON labels_data.uuid = labels.uuid
  LEFT JOIN labels_deletes on labels.uuid = labels_deletes.uuid
WHERE labels_deletes.uuid IS NULL
ORDER BY labels_data.created_at DESC
LIMIT 1;

-- name: UpdateLabel :exec
INSERT INTO labels_data (uuid, customer, family, model, part_number, station, label)
  VALUES(?, ?, ?, ?, ?, ?, ?);

-- name: DeleteLabel :exec
INSERT INTO labels_deletes (uuid) VALUES (?);
