-- name: CreateMeasuredVariable :one
INSERT INTO measuredvariables (
  variable,
  alias,
  description
) VALUES(
  $1, $2, $3
) RETURNING *;

-- name: GetMeasuredVariable :one
SELECT * FROM measuredvariables
WHERE id = $1 LIMIT 1;

-- name: ListMeasuredVariables :many
SELECT * FROM measuredvariables
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteMeasuredVariable :exec
DELETE FROM measuredvariables
WHERE id = $1;