-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NOT NULL DEFAULT 1,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);

-- Table to track changes on roles
CREATE TABLE IF NOT EXISTS archived_roles (
    id UUID PRIMARY KEY NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NOT NULL DEFAULT 1,

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (id) REFERENCES roles (id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS archived_roles;
DROP TABLE IF EXISTS roles;