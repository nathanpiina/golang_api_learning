-- +goose Up
-- +goose StatementBegin
CREATE TABLE People (
    Id SERIAL PRIMARY KEY,
    nickname VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100),
    birth VARCHAR(20),
    stack VARCHAR(50)
);

-- +goose Down
DROP TABLE People;