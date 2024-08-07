-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id          SERIAL PRIMARY KEY,
    first_name  varchar(255),
    middle_name varchar(255),
    last_name   varchar(255),
    email       varchar(255) UNIQUE,
    phone       varchar(255),
    avatar      varchar(255),
    about       text,
    theme       varchar(255),
    lang        varchar(255) default 'ru',
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
