BEGIN;

-- Enable uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create schema for API on the database
CREATE SCHEMA genesis_server;

COMMIT;
