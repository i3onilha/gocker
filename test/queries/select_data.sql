SELECT
  labels_data.id,
  labels_data.customer,
  labels_data.model,
  labels_data.part_number,
  labels_data.station,
  labels_data.dpi
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
LIMIT 10 OFFSET 0;
