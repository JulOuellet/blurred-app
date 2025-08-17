CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS sports (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    name text NOT NULL
);

INSERT INTO sports (name)
    VALUES ('Road Cycling');

