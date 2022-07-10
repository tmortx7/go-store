-- name: CreateInstrumentType :one
INSERT INTO instrumenttypes(
  type,
  alias,
  description
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetInstrumentType :one
SELECT * FROM instrumenttypes
WHERE id = $1 LIMIT 1;

-- name: ListInstrumentTypes :many
SELECT * FROM instrumenttypes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteInstrumentType :exec
DELETE FROM instrumenttypes
WHERE id = $1;
