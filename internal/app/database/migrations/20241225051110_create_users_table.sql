-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    username VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(60) NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Seed data
INSERT INTO users (id, username, email, password_hash) VALUES
    ('00000000-0000-0000-0000-000000000001', 'adminusername', 'admin@email.com', '$2a$10$AMLoK2T8XDVO2XAX0maDJeFpPwQfCTkzdRNceW9pp.f2saSIYP9Me'),
    ('00000000-0000-0000-0000-000000000002', 'superadminusername', 'superadmin@email.com', '$2a$10$AMLoK2T8XDVO2XAX0maDJeFpPwQfCTkzdRNceW9pp.f2saSIYP9Me');

-- Table to track changes on users
CREATE TABLE IF NOT EXISTS archived_users (
    id UUID PRIMARY KEY NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(60) NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_users;
DROP TABLE IF EXISTS users;
