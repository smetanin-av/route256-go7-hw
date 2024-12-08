-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages
(
    message_id uuid      NOT NULL UNIQUE,
    created_at timestamp NOT NULL,
    order_id   int8      NOT NULL,
    status_old varchar   NOT NULL,
    status_new varchar   NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
