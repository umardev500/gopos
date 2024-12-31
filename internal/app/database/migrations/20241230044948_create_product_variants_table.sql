-- +goose Up
CREATE TABLE IF NOT EXISTS product_variants (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL, -- Variant name (e.g., "Small", "Medium", "Large")
    description TEXT,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Table to track changes on product variants
CREATE TABLE IF NOT EXISTS archived_product_variants (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    product_id UUID NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL, -- Variant name (e.g., "Small", "Medium", "Large")
    description TEXT,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change
    
    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_product_variants;
DROP TABLE IF EXISTS product_variants;
