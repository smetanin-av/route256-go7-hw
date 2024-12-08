-- +goose Up
-- +goose StatementBegin
CREATE TABLE carts
(
    user_id  int8 NOT NULL,
    sku   int4 NOT NULL,
    quantity int2 NOT NULL,
    UNIQUE (user_id, sku)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE carts;
-- +goose StatementEnd
