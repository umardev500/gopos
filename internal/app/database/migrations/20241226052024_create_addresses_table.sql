-- +goose Up
CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    country_id INT NOT NULL,
    province_id INT NOT NULL,
    city_id INT NOT NULL,
    "address" TEXT NULL DEFAULT NULL,
    zip_code VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    deleted_at TIMESTAMPTZ NULL DEFAULT NULL
);

-- +goose Down
DROP TABLE IF EXISTS addresses;
