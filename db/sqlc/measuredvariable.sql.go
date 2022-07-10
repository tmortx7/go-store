// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: measuredvariable.sql

package db

import (
	"context"
	"database/sql"
)

const createMeasuredVariable = `-- name: CreateMeasuredVariable :one
INSERT INTO measuredvariables (
  variable,
  alias,
  description
) VALUES(
  $1, $2, $3
) RETURNING id, variable, alias, description, created_at
`

type CreateMeasuredVariableParams struct {
	Variable    string         `json:"variable"`
	Alias       string         `json:"alias"`
	Description sql.NullString `json:"description"`
}

func (q *Queries) CreateMeasuredVariable(ctx context.Context, arg CreateMeasuredVariableParams) (Measuredvariable, error) {
	row := q.db.QueryRowContext(ctx, createMeasuredVariable, arg.Variable, arg.Alias, arg.Description)
	var i Measuredvariable
	err := row.Scan(
		&i.ID,
		&i.Variable,
		&i.Alias,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteMeasuredVariable = `-- name: DeleteMeasuredVariable :exec
DELETE FROM measuredvariables
WHERE id = $1
`

func (q *Queries) DeleteMeasuredVariable(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMeasuredVariable, id)
	return err
}

const getMeasuredVariable = `-- name: GetMeasuredVariable :one
SELECT id, variable, alias, description, created_at FROM measuredvariables
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMeasuredVariable(ctx context.Context, id int64) (Measuredvariable, error) {
	row := q.db.QueryRowContext(ctx, getMeasuredVariable, id)
	var i Measuredvariable
	err := row.Scan(
		&i.ID,
		&i.Variable,
		&i.Alias,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listMeasuredVariables = `-- name: ListMeasuredVariables :many
SELECT id, variable, alias, description, created_at FROM measuredvariables
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListMeasuredVariablesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListMeasuredVariables(ctx context.Context, arg ListMeasuredVariablesParams) ([]Measuredvariable, error) {
	rows, err := q.db.QueryContext(ctx, listMeasuredVariables, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Measuredvariable{}
	for rows.Next() {
		var i Measuredvariable
		if err := rows.Scan(
			&i.ID,
			&i.Variable,
			&i.Alias,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}