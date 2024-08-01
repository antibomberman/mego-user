-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id       BIGINT PRIMARY KEY AUTO_INCREMENT,
    first_name varchar(255),
    middle_name varchar(255),
    last_name  varchar(255),
    email      varchar(255) UNIQUE,
    phone varchar(255),
    avatar varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by varchar(255),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by varchar(255),
    deleted_at TIMESTAMP ,
    deleted_by varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
