// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1
// source: query.sql

package querytest

import (
	"context"
)

const identicalTable = `-- name: IdenticalTable :many
SELECT id FROM foo
`

func (q *Queries) IdenticalTable(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, identicalTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
