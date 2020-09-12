BEGIN;

CREATE TABLE rooms (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR (255) NOT NULL,
    status VARCHAR (255) NOT NULL,
    team_id uuid NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    FOREIGN KEY (team_id) REFERENCES teams (id)
);

COMMIT;