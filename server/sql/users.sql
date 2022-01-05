-- name: CreateUser :exec
INSERT INTO genesis_server.users
    (email, password_hash, first_name)
VALUES
    ($1, $2, $3);

-- name: GetUserByID :one
SELECT * FROM genesis_server.users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM genesis_server.users
WHERE email = $1;

-- name: UpdateUserPassword :exec
UPDATE genesis_server.users
SET password_hash = $1, password_uuid = uuid_generate_v4()
WHERE id = $1;
