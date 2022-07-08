-- name: CreateProcessVariable :one
INSERT INTO processvariables (
  name,
  alias,
  description
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProcessVariable :one
SELECT * FROM processvariables
WHERE id = $1 LIMIT 1;

-- name: ListProcessVariables :many
SELECT * FROM processvariables
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteProcessVariable :exec
DELETE FROM processvariables
WHERE id = $1;