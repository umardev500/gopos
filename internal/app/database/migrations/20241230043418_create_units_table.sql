-- +goose Up
CREATE TABLE IF NOT EXISTS units (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NOT NULL DEFAULT 1,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    UNIQUE (tenant_id, name),
    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE
);

-- Table to track changes on product units
CREATE TABLE IF NOT EXISTS archived_units (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NOT NULL DEFAULT 1,

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS archived_units;
DROP TABLE IF EXISTS units;
