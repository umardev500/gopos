-- +goose Up
CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Table to track changes on user roles
CREATE TABLE IF NOT EXISTS archived_user_roles (
    user_id UUID NOT NULL,
    role_id UUID NOT NULL,
    version INT NULL,

    modified_by UUID NULL, -- User who made the change
    
    archived_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL,

   PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE,
    FOREIGN KEY (modified_by) REFERENCES users (id) ON DELETE SET NULL
);

-- +goose Down
DROP TABLE IF EXISTS archived_user_roles;
DROP TABLE IF EXISTS user_roles;
