-- +goose Up
CREATE TABLE IF NOT EXISTS user_branches (
    user_id UUID NOT NULL,
    branch_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,
    
    PRIMARY KEY (user_id, branch_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_branches;
