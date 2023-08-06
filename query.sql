SELECT
  labels_data.*
FROM
  labels_data
  LEFT JOIN labels_deletes ON labels_data.id = labels_deletes.id
WHERE labels_deletes.id IS NULL
  AND labels_data.id = 214 
ORDER BY labels_data.created_at DESC
LIMIT 1;
