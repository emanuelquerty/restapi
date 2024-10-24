-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY NOT NULL,
    first_name VARCHAR (20),
    last_name VARCHAR (20),
    email VARCHAR (255) UNIQUE,
    password VARCHAR (64) UNIQUE

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
