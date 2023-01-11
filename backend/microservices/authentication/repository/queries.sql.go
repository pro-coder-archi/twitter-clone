// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package repository

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users
    (email, password)
        VALUES ($1, $2)
`

type CreateUserParams struct {
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser, arg.Email, arg.Password)
	return err
}

const findRegisteredEmail = `-- name: FindRegisteredEmail :one
SELECT id, email, password FROM users
    WHERE users.email= $1
        LIMIT 1
`

func (q *Queries) FindRegisteredEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, findRegisteredEmail, email)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return i, err
}
