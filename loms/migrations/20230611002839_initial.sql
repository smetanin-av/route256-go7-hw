-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    order_id bigserial PRIMARY KEY,
    user_id  int8 NOT NULL,
    status   int4 NOT NULL
);

CREATE TABLE items
(
    order_id bigint NOT NULL,
    sku      int4   NOT NULL,
    quantity int2   NOT NULL
);

CREATE TABLE stocks
(
    sku          int4,
    warehouse_id int8,
    quantity     int8,
    UNIQUE (sku, warehouse_id)
);

CREATE TABLE reserved
(
    order_id     int8,
    sku          int4,
    warehouse_id int8,
    quantity     int8
);

CREATE TABLE sold
(
    order_id     int8,
    sku          int4,
    warehouse_id int8,
    quantity     int8
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;

DROP TABLE items;

DROP TABLE stocks;

DROP TABLE reserved;

DROP TABLE sold;
-- +goose StatementEnd
