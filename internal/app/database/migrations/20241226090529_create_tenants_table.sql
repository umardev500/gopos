-- +goose Up
CREATE TABLE IF NOT EXISTS tenants (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    address_id UUID NULL DEFAULT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,

    modified_by UUID NULL, -- User who made the change

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (address_id) REFERENCES addresses (id) ON DELETE SET NULL,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Archived tenants
CREATE TABLE IF NOT EXISTS archived_tenants (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL,
    username VARCHAR(25) UNIQUE NOT NULL,
    address_id UUID NULL DEFAULT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,

    modified_by UUID NULL, -- User who made the change

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (address_id) REFERENCES addresses (id) ON DELETE SET NULL,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_tenants;
DROP TABLE IF EXISTS tenants;
