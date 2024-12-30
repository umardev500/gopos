-- +goose Up
CREATE TABLE IF NOT EXISTS branch_product_inventory (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    branch_id UUID NOT NULL,
    product_id UUID NOT NULL, -- Links to shared prodcut
    unit_id UUID NOT NULL,
    variant_id UUID NULL, -- Optional: Link to specific variant
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Product/unit availability in branch
    version INT NOT NULL DEFAULT 1,

    user_id UUID NOT NULL, -- User who made the change

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES product_variants (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

-- Table to track changes on branch product inventory
CREATE TABLE IF NOT EXISTS archived_branch_product_inventory (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    branch_id UUID NOT NULL,
    product_id UUID NOT NULL, -- Links to shared prodcut
    unit_id UUID NOT NULL,
    variant_id UUID NULL, -- Optional: Link to specific variant
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Product/unit availability in branch
    version INT NOT NULL DEFAULT 1,

    user_id UUID NOT NULL, -- User who made the change

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (unit_id) REFERENCES units (id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES product_variants (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_branch_product_inventory;
DROP TABLE IF EXISTS branch_product_inventory;
