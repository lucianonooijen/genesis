-- name: AddPasswordForgotForUser :exec
-- Adds password_forgot entry for user.
-- In case there is a password reset (database unique conflict)
-- the existing entry will be reset with new values
INSERT INTO password_forgot(user_id)
VALUES ($1)
ON CONFLICT (user_id)
DO UPDATE SET
    reset_token = DEFAULT,
    valid_until = DEFAULT,
    is_used = DEFAULT;

-- name: GetPasswordResetByUserId :one
SELECT * FROM password_forgot
WHERE user_id = $1;

-- name: GetPasswordResetByResetToken :one
SELECT * FROM password_forgot
WHERE reset_token = $1;

-- name: MarkResetTokenUsed :exec
-- Marks a password_forgot entry used based on user_id
UPDATE password_forgot
SET is_used = true
WHERE user_id = $1;
