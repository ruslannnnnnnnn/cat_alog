CREATE TABLE IF NOT EXISTS catalog.cats
(
    id            UUID PRIMARY KEY,
    name          TEXT,
    date_of_birth DATE,
    image_url     TEXT
);