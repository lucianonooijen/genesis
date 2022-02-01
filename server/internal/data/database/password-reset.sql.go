// Code generated by sqlc. DO NOT EDIT.
// source: password-reset.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const addPasswordForgotForUser = `-- name: AddPasswordForgotForUser :exec
INSERT INTO password_forgot(user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE SET
    reset_token = DEFAULT,
    valid_until = DEFAULT,
    is_used = DEFAULT
`

// Adds password_forgot entry for user.
// In case there is a password reset (database unique conflict)
// the existing entry will be reset with new values
func (q *Queries) AddPasswordForgotForUser(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, addPasswordForgotForUser, userID)
	return err
}

const deletePasswordForgotByUserId = `-- name: DeletePasswordForgotByUserId :exec
DELETE FROM password_forgot
WHERE user_id = $1
`

// Deletes password_forgot for user
func (q *Queries) DeletePasswordForgotByUserId(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deletePasswordForgotByUserId, userID)
	return err
}

const getPasswordResetByResetToken = `-- name: GetPasswordResetByResetToken :one
SELECT user_id, reset_token, valid_until, is_used FROM password_forgot
WHERE reset_token = $1
`

func (q *Queries) GetPasswordResetByResetToken(ctx context.Context, resetToken uuid.UUID) (PasswordForgot, error) {
	row := q.db.QueryRowContext(ctx, getPasswordResetByResetToken, resetToken)
	var i PasswordForgot
	err := row.Scan(
		&i.UserID,
		&i.ResetToken,
		&i.ValidUntil,
		&i.IsUsed,
	)
	return i, err
}

const getPasswordResetByUserId = `-- name: GetPasswordResetByUserId :one
SELECT user_id, reset_token, valid_until, is_used FROM password_forgot
WHERE user_id = $1
`

func (q *Queries) GetPasswordResetByUserId(ctx context.Context, userID int32) (PasswordForgot, error) {
	row := q.db.QueryRowContext(ctx, getPasswordResetByUserId, userID)
	var i PasswordForgot
	err := row.Scan(
		&i.UserID,
		&i.ResetToken,
		&i.ValidUntil,
		&i.IsUsed,
	)
	return i, err
}

const markResetTokenUsed = `-- name: MarkResetTokenUsed :exec
UPDATE password_forgot
SET is_used = true
WHERE reset_token = $1
`

// Marks a password_forgot entry used based on user_id
func (q *Queries) MarkResetTokenUsed(ctx context.Context, resetToken uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, markResetTokenUsed, resetToken)
	return err
}
