-- +goose Up
-- +goose StatementBegin
CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS books
(
    ID          UUID primary key DEFAULT uuid_generate_v4(),
    CREATED_AT  TIMESTAMP NOT NULL,
    UPDATED_AT  TIMESTAMP NOT NULL,
    NAME        VARCHAR   NOT NULL,
    AUTHOR      VARCHAR   NOT NULL,
    ISBN        VARCHAR   NOT NULL,
    DESCRIPTION VARCHAR
);

CREATE TABLE IF NOT EXISTS borrows
(
    ID           UUID primary key DEFAULT uuid_generate_v4(),
    CREATED_AT   TIMESTAMP                  NOT NULL,
    UPDATED_AT   TIMESTAMP                  NOT NULL,
    BOOK_ID      UUID references books (id) NOT NULL,
    TAKEN_DATE   TIMESTAMP                  NOT NULL,
    BROUGHT_DATE TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
