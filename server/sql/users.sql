-- name: CreateUser :exec
-- Inserts new user into database
INSERT INTO users
    (email, password_hash, first_name)
VALUES
    ($1, $2, $3);

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByUuid :one
SELECT * FROM users
WHERE user_uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserPassword :exec
-- Update password_hash for user based on id.
-- Also resets the password_uuid to invalidate existing JWTs
UPDATE users
SET password_hash = $1, password_uuid = uuid_generate_v4()
WHERE id = $2;

-- name: GetUserProfileByID :one
-- Gets the user profile based on user profile id
SELECT first_name
FROM users
WHERE id = $1;

-- name: UpdateUserProfile :exec
-- Updates user profile
UPDATE users
SET first_name = $1
WHERE id = $2;

-- name: DeleteUser :exec
-- Deletes the user profile
DELETE FROM users
WHERE id = $1;
