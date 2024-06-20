GOPATH:=$(shell go env GOPATH)

CONFIG_FILE:=config.yml

.PHONY: generate
generate: generate_client
	rm -rf pkg/openapi3/models
	openapi-generator generate -i pkg/openapi3/openapi3.yaml -o pkg/openapi3 -g go-echo-server
	rm -rf internal/entities/models && cp -r pkg/openapi3/models internal/entities/models;

.PHONY: generate_client
generate_client:
	rm -rf pkg/openapi3/client
	openapi-generator generate -i pkg/openapi3/openapi3.yaml -o pkg/openapi3/client -g go
	rm -rf internal/itest/client && rsync -a --exclude='go.mod' --exclude='go.sum' pkg/openapi3/client internal/itest;

.PHONY: build
build: generate
	go build -o ./build/rest cmd/rest_server/main.go
	go build -o ./build/puller cmd/subgraph_puller/main.go

.PHONY: init_db
init_db:
	go get -u -d github.com/mattes/migrate/cli
	go get -u -d github.com/go-sql-driver/mysql
	go build -tags 'mysql' -o /usr/local/bin/migrate github.com/mattes/migrate/cli
	go run tools/reset_sight/main.go

.PHONY: rest_docker
rest_docker: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/rest cmd/rest_server/main.go
	cd build; docker build . -f Dockerfile_rest -t rest:latest --build-arg config_file=$(CONFIG_FILE);

.PHONY: puller_docker
puller_docker: generate
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/puller cmd/subgraph_puller/main.go
	cd build; docker build . -f Dockerfile_puller -t puller:latest --build-arg config_file=$(CONFIG_FILE);

.PHONY: run_rest
run_rest: build
	./build/rest -config-path ./build/  -config-name config

.PHONY: run_puller
run_puller: build
	./build/puller -config-path ./build/  -config-name config
