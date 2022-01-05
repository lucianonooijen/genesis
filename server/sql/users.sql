-- name: CreateUser :exec
-- Inserts new user into database
INSERT INTO users
    (email, password_hash, first_name)
VALUES
    ($1, $2, $3);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET password_hash = $1, password_uuid = uuid_generate_v4()
WHERE id = $1;
