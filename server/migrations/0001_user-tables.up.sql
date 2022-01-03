BEGIN;

-- Init users table
CREATE TABLE genesis_server.users (
    id SERIAL PRIMARY KEY,
    user_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    password_hash varchar NOT NULL,
    password_uuid uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),
    email varchar UNIQUE NOT NULL,
    first_name varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now())
);

-- Init password forgot table
CREATE TABLE genesis_server.password_forgot (
    user_id int UNIQUE NOT NULL REFERENCES genesis_server.users(id),
    reset_token uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),  -- Must be unique, to avoid the possibility of setting other user's passwords
    valid_until timestamp DEFAULT (now() + '1 day'::interval) NOT NULL,
    is_used boolean DEFAULT false NOT NULL
);

COMMIT;
