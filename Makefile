.PHONY: generate
generate:
	protoc \
		--proto_path=api/proto/v1 \
		--proto_path=/usr/local/include \
		--go_out=api/proto/v1 \
		--go_opt=paths=source_relative \
		--go-grpc_out=api/proto/v1 \
		--go-grpc_opt=paths=source_relative \
		api/proto/v1/*.proto

.PHONY: build
build: generate
	go build -o bin/server cmd/server/main.go

.PHONY: run
run: build
	./bin/server
