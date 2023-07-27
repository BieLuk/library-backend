-- +goose Up
-- +goose StatementBegin
CREATE
    EXTENSION "uuid-ossp";

CREATE TABLE authors
(
    ID         UUID primary key,
    CREATED_AT TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP NOT NULL,
    NAME       VARCHAR   NOT NULL,
    SURNAME    VARCHAR   NOT NULL
);

CREATE TABLE borrowers
(
    ID         UUID primary key,
    CREATED_AT TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP NOT NULL,
    NAME       VARCHAR   NOT NULL,
    SURNAME    VARCHAR   NOT NULL
);

CREATE TABLE books
(
    ID          UUID primary key,
    CREATED_AT  TIMESTAMP                    NOT NULL,
    UPDATED_AT  TIMESTAMP                    NOT NULL,
    NAME        VARCHAR                      NOT NULL,
    AUTHOR_ID   UUID references authors (id) NOT NULL,
    ISBN        VARCHAR                      NOT NULL,
    DESCRIPTION VARCHAR
);

CREATE TABLE borrows
(
    ID         UUID primary key,
    CREATED_AT TIMESTAMP NOT NULL,
    UPDATED_AT TIMESTAMP NOT NULL,
    BORROWER_ID UUID references borrowers(id) NOT NULL,
    BOOK_ID UUID references books(id) NOT NULL,
    TAKEN_DATE TIMESTAMP NOT NULL,
    BROUGHT_DATE TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
