-- +goose Up
CREATE TABLE IF NOT EXISTS user_branches (
    user_id UUID NOT NULL,
    branch_id UUID NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    PRIMARY KEY (user_id, branch_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Seed data
INSERT INTO user_branches (user_id, branch_id) VALUES
    ('00000000-0000-0000-0000-000000000001', '00000000-0000-0000-0000-000000000001');

-- Table to track changes on product units
CREATE TABLE IF NOT EXISTS archived_user_branches (
    user_id UUID NOT NULL,
    branch_id UUID NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change

    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    PRIMARY KEY (user_id, branch_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_user_branches;
DROP TABLE IF EXISTS user_branches;
