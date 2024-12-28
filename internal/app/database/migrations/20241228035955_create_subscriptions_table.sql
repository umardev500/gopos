-- +goose Up
CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    plan_id UUID NOT NULL,
    start_date DATE NOT NULL, -- Start of the subscription
    end_date DATE NOT NULL, -- End of the subscription
    is_active BOOLEAN NOT NULL DEFAULT TRUE, -- Tracks whether the subscription is active (TRUE/FALSE)
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    FOREIGN KEY (tenant_id) REFERENCES tenants(id) ON DELETE CASCADE,
    FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;