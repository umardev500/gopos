-- +goose Up
CREATE TABLE IF NOT EXISTS countries (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Seed dummy data
INSERT INTO countries (id, name) VALUES
    (1, 'Indonesia');

CREATE TABLE IF NOT EXISTS provinces (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country_id INT NOT NULL,
    FOREIGN KEY (country_id) REFERENCES countries (id)
);

-- Seed dummy data
INSERT INTO provinces (id, name, country_id) VALUES
    (1, 'Banten', 1),
    (2, 'Jawa Tengah', 1),
    (3, 'Jawa Timur', 1);

-- Kabupaten
CREATE TABLE IF NOT EXISTS cities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    province_id INT NOT NULL,
    FOREIGN KEY (province_id) REFERENCES provinces (id)
);

-- Seed dummy data
INSERT INTO cities (id, name, province_id) VALUES
    (1, 'Pandeglang', 1),
    (2, 'Bandung', 2),
    (3, 'Semarang', 3);


-- Kecamatan
CREATE TABLE IF NOT EXISTS districts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    city_id INT NOT NULL,
    FOREIGN KEY (city_id) REFERENCES cities (id)
);

-- Seed dummy data
INSERT INTO districts (id, name, city_id) VALUES
    (1, 'Pandeglang', 1),
    (2, 'Bandung', 2),
    (3, 'Semarang', 3);


-- +goose Down
DROP TABLE IF EXISTS districts;
DROP TABLE IF EXISTS cities;
DROP TABLE IF EXISTS provinces;
DROP TABLE IF EXISTS countries;

