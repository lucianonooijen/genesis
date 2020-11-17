BEGIN;

CREATE TABLE password_forgot (
    user_id int UNIQUE NOT NULL REFERENCES users(id),
    reset_token uuid UNIQUE NOT NULL DEFAULT (uuid_generate_v4()),  -- Must be unique, to avoid the possibility of setting other user's passwords
    valid_until timestamp DEFAULT (now() + '1 day'::interval) NOT NULL,
    is_used boolean DEFAULT false NOT NULL
);

COMMIT;
