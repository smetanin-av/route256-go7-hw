CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
PACKAGE=route256/loms/cmd/app

.PHONY: all
all: format build test lint

.PHONY: build
build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

.PHONY: test
test:
	go test ./...

.PHONY: run
run:
	go run ${PACKAGE}

.PHONY: lint
lint: install-lint
	${LINTBIN} run

.PHONY: precommit
precommit: format build test lint
	echo "OK"

.PHONY: bindir
bindir:
	mkdir -p ${BINDIR}

.PHONY: format
format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

.PHONY: install-lint
install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

.PHONY: install-smartimports
install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

.PHONY: install-protoc
install-protoc: bindir
	GOBIN=$(BINDIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(BINDIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(BINDIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(BINDIR) go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: generate
generate:
	PATH="$$PATH:$(BINDIR)" \
	GO_OUT="pkg" \
	GO_OPT="paths=source_relative" && \
	protoc -I api/ -I vendor.proto/ \
	--go_out $$GO_OUT --go_opt $$GO_OPT \
	--go-grpc_out $$GO_OUT --go-grpc_opt $$GO_OPT \
	--grpc-gateway_out $$GO_OUT --grpc-gateway_opt $$GO_OPT \
	--validate_out="lang=go,$$GO_OPT:$$GO_OUT" \
	$(shell find api/ -iname "*.proto")
