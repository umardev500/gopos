-- +goose Up
CREATE TABLE IF NOT EXISTS plans (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    highlights JSONB NOT NULL,
    metadata JSONB NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    version INT NOT NULL DEFAULT 1,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);


-- Table to track changes on subsriptions
CREATE TABLE IF NOT EXISTS archived_plans (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    plan_id UUID NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    highlights JSONB NOT NULL,
    metadata JSONB NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    version INT NOT NULL DEFAULT 1,

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_plans;
DROP TABLE IF EXISTS plans;
