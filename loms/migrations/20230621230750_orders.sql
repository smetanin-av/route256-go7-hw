-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
    ADD COLUMN created_at timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
    DROP COLUMN created_at;
-- +goose StatementEnd
