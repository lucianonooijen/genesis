-- name: GetPushTokensForUser :many
-- Gives all push notification tokens for given user
SELECT platform, token
FROM user_push_tokens
WHERE userid = $1;


-- name: AddUserPushToken :exec
-- Adds push token to the database for a given user and platform
INSERT INTO user_push_tokens (userid, platform, token)
VALUES ($1, $2, $3)
ON CONFLICT DO NOTHING;
