-- +goose Up
CREATE TABLE IF NOT EXISTS user_tenants (
    user_id UUID NOT NULL,
    tenant_id UUID NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    PRIMARY KEY (user_id, tenant_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (tenant_id) REFERENCES tenants (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_tenants;
