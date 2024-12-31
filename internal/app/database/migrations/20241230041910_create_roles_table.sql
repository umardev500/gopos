-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NULL, -- Nullable for platform roles
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Seed data
INSERT INTO roles (id, tenant_id, name, description) VALUES
    ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001', 'admin', 'Admin role'),
    ('00000000-0000-0000-0000-000000000002', '00000000-0000-0000-0000-000000000001', 'user', 'User role');

-- Table to track changes on roles
CREATE TABLE IF NOT EXISTS archived_roles (
    id UUID PRIMARY KEY NOT NULL,
    tenant_id UUID NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (id) REFERENCES roles (id) ON DELETE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);


-- +goose Down
DROP TABLE IF EXISTS archived_roles;
DROP TABLE IF EXISTS roles;