DC_CFG_ENV=-f docker-compose-env.yml
DC_CFG_APPS=${DC_CFG_ENV} -f docker-compose-app.yml
DC_SRV_APPS=checkout_app loms_app notific_app

.PHONY: build-all
build-all:
	cd checkout && GOOS=linux GOARCH=amd64 make build
	cd loms && GOOS=linux GOARCH=amd64 make build
	cd notifications && GOOS=linux GOARCH=amd64 make build

.PHONY: precommit
precommit:
	cd libs && make precommit
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

.PHONY: fix-format
fix-format:
	gci write --skip-generated -s standard,default .

.PHONY: dc-up-env
dc-up-env:
	docker-compose ${DC_CFG_ENV} up -d

.PHONY: dc-down-env
dc-down-env:
	docker-compose ${DC_CFG_ENV} down --remove-orphans --volumes

.PHONY: dc-up-apps
dc-up-apps: build-all
	docker-compose ${DC_CFG_APPS} up --force-recreate --build ${DC_SRV_APPS} -d

.PHONY: dc-stop-apps
dc-stop-apps:
	docker-compose ${DC_CFG_APPS} stop --timeout=10 ${DC_SRV_APPS}

.PHONY: dc-up-all
dc-up-all: build-all
	docker-compose ${DC_CFG_APPS} up --force-recreate --build -d

.PHONY: dc-down-all
dc-down-all:
	docker-compose ${DC_CFG_APPS} down --remove-orphans --volumes

.PHONY: goose-status
goose-status:
	CHECKOUT_DSN="postgres://db_user:db_password@localhost:15432/checkout?sslmode=disable" && \
	goose -dir ./checkout/migrations postgres $$CHECKOUT_DSN status

	LOMS_DSN="postgres://db_user:db_password@localhost:25432/loms?sslmode=disable" && \
	goose -dir ./loms/migrations postgres $$LOMS_DSN status

	NOTIFIC_DSN="postgres://db_user:db_password@notific_pgb:35432/notific?sslmode=disable" && \
	goose -dir ./notifications/migrations postgres $$NOTIFIC_DSN status

.PHONY: goose-up
goose-up:
	CHECKOUT_DSN="postgres://db_user:db_password@localhost:15432/checkout?sslmode=disable" && \
	goose -dir ./checkout/migrations postgres $$CHECKOUT_DSN up

	LOMS_DSN="postgres://db_user:db_password@localhost:25432/loms?sslmode=disable" && \
	goose -dir ./loms/migrations postgres $$LOMS_DSN up

	NOTIFIC_DSN="postgres://db_user:db_password@localhost:35432/notific?sslmode=disable" && \
	goose -dir ./notifications/migrations postgres $$NOTIFIC_DSN up

.PHONY: goose-down
goose-down:
	CHECKOUT_DSN="postgres://db_user:db_password@localhost:15432/checkout?sslmode=disable" && \
	goose -dir ./checkout/migrations postgres $$CHECKOUT_DSN down

	LOMS_DSN="postgres://db_user:db_password@localhost:25432/loms?sslmode=disable" && \
	goose -dir ./loms/migrations postgres $$LOMS_DSN down

	NOTIFIC_DSN="postgres://db_user:db_password@localhost:35432/notific?sslmode=disable" && \
	goose -dir ./notifications/migrations postgres $$NOTIFIC_DSN down

.PHONY: e2e-tests
e2e-tests: dc-up-all goose-up
	sleep 5

	curl "localhost:28080/v1/stocks/1076963" | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 301}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 1}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq
	curl "localhost:18080/v1/deleteFromCart" -d '{"user": 1, "sku": 1076963, "count": 1}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 1}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1148162, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1625903, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 2618151, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 2956315, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 3596599, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 3618852, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 4288068, "count": 1}' | jq
	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 4465995, "count": 1}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq

	curl -v "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 1}'
	curl "localhost:18080/v1/purchase" -d '{"user": 1}' | jq
	curl "localhost:18080/v1/listCart" -d '{"user": 1}' | jq
	curl "localhost:28080/v1/stocks/1076963" | jq
	curl "localhost:28080/v1/listOrder" -d '{"order_id": 1}' | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 1}' | jq
	curl "localhost:18080/v1/purchase" -d '{"user": 1}' | jq
	curl "localhost:28080/v1/orderPayed" -d '{"order_id": 2}' | jq
	curl "localhost:28080/v1/listOrder" -d '{"order_id": 2}' | jq
	curl "localhost:28080/v1/stocks/1076963" | jq

	curl "localhost:18080/v1/addToCart" -d '{"user": 1, "sku": 1076963, "count": 75}' | jq
	curl "localhost:18080/v1/purchase" -d '{"user": 1}' | jq
	curl "localhost:28080/v1/cancelOrder" -d '{"order_id": 3}' | jq
	curl "localhost:28080/v1/listOrder" -d '{"order_id": 3}' | jq
	curl "localhost:28080/v1/stocks/1076963" | jq

.PHONY: unit-tests
unit-tests:
	go test -v -race -coverprofile coverage.out ./checkout/internal/... ./libs/... ./loms/internal/...
	go tool cover -func=coverage.out
