-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
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

-- Table to track changes on products
CREATE TABLE IF NOT EXISTS archived_products (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change
    
    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE,
    FOREIGN KEY (id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_products;
DROP TABLE IF EXISTS products;
