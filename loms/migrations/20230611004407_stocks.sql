-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks(sku, warehouse_id, quantity)
VALUES (1076963, 1, 100),
       (1076963, 2, 200),
       (1148162, 1, 10),
       (1148162, 2, 2),
       (1625903, 1, 1000),
       (1625903, 2, 10),
       (2618151, 1, 10),
       (2618151, 2, 10),
       (2956315, 2, 100),
       (2958025, 2, 56),
       (3596599, 1, 89),
       (3618852, 2, 45),
       (4288068, 2, 90),
       (4465995, 1, 12);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM stocks
WHERE sku IN (1076963, 1148162, 1625903, 2618151, 2956315, 2958025, 3596599, 3618852, 4288068, 4465995)
  AND warehouse_id IN (1, 2);
-- +goose StatementEnd
