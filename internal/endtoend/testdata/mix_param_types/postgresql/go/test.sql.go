// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.1
// source: test.sql

package querytest

import (
	"context"
)

const countFour = `-- name: CountFour :one
SELECT count(1) FROM bar WHERE id > ? AND phone <> ? AND name <> ?
`

type CountFourParams struct {
	ID    int32
	Phone string
	Name  string
}

func (q *Queries) CountFour(ctx context.Context, arg CountFourParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countFour, arg.ID, arg.Phone, arg.Name)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countOne = `-- name: CountOne :one
SELECT count(1) FROM bar WHERE id = $2 AND name <> $1 LIMIT $3
`

type CountOneParams struct {
	Name  string
	ID    int32
	Limit int32
}

func (q *Queries) CountOne(ctx context.Context, arg CountOneParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countOne, arg.Name, arg.ID, arg.Limit)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countThree = `-- name: CountThree :one
SELECT count(1) FROM bar WHERE id > $2 AND phone <> $3 AND name <> $1
`

type CountThreeParams struct {
	Name  string
	ID    int32
	Phone string
}

func (q *Queries) CountThree(ctx context.Context, arg CountThreeParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countThree, arg.Name, arg.ID, arg.Phone)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countTwo = `-- name: CountTwo :one
SELECT count(1) FROM bar WHERE id = $1 AND name <> $2
`

type CountTwoParams struct {
	ID   int32
	Name string
}

func (q *Queries) CountTwo(ctx context.Context, arg CountTwoParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, countTwo, arg.ID, arg.Name)
	var count int64
	err := row.Scan(&count)
	return count, err
}
