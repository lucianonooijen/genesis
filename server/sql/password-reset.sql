-- name: AddPasswordForgotForUser :exec
INSERT INTO genesis_server.password_forgot(user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE SET
    reset_token = DEFAULT,
    valid_until = DEFAULT,
    is_used = DEFAULT;

-- name: GetPasswordResetByUserId :one
SELECT * FROM genesis_server.password_forgot
WHERE user_id = $1;

-- name: GetPasswordResetByResetToken :one
SELECT * FROM genesis_server.password_forgot
WHERE reset_token = $1;

-- name: MarkResetTokenUsed :exec
UPDATE genesis_server.password_forgot
SET is_used = true
WHERE user_id = $1;
