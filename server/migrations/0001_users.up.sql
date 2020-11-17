BEGIN;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    password_hash varchar NOT NULL,
    password_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    email varchar UNIQUE NOT NULL,
    first_name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now())
);

COMMIT;
