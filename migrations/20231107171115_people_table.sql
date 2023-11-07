-- +goose Up
-- +goose StatementBegin
CREATE TABLE People (
    nickname VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100),
    birth DATE,
    stack VARCHAR(100)
);
-- +goose Down
DROP TABLE People;