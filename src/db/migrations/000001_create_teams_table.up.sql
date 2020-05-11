BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE teams (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR (255) NOT NULL,
    status VARCHAR (255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

COMMIT;