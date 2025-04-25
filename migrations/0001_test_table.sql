-- +goose Up
CREATE TABLE IF NOT EXISTS test_table (
    id SERIAL PRIMARY KEY,
    test_name TEXT NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS test_table;

