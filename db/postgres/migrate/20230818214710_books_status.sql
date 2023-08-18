-- +goose Up
-- +goose StatementBegin
ALTER TABLE books
    ADD COLUMN STATUS VARCHAR NOT NULL DEFAULT 'ACTIVE';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
