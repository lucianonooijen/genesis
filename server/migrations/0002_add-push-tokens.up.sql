BEGIN;

CREATE TYPE mobile_platform AS ENUM ('Android', 'iOS');

CREATE TABLE user_push_tokens (
    userid int NOT NULL REFERENCES users(id),
    platform mobile_platform NOT NULL,
    token varchar NOT NULL UNIQUE
);

COMMIT;
