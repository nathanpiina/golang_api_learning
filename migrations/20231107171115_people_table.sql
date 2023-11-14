-- +goose Up
-- +goose StatementBegin
CREATE TABLE People (
    Id INT PRIMARY KEY AUTO_INCREMENT,
    nickname VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(100),
    birth DATE,
    stack VARCHAR(50)
);

-- +goose Down
DROP TABLE People;