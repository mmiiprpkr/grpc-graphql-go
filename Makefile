build:
	go build -o bin/server .

run: build
	./bin/server

# __________________________ GENERATED SECTION __________________________
clear-generated:
	rm -rf generated

proto-gen:
	mkdir -p generated/proto && \
	protoc --go_out=./generated/proto --go_opt=paths=source_relative \
    --go-grpc_out=./generated/proto --go-grpc_opt=paths=source_relative \
    --proto_path=proto \
    proto/**/*.proto --experimental_allow_proto3_optional

gql-gen:
	mkdir -p generated && go run github.com/99designs/gqlgen run

enum-gen:
	enum-gen

gen: clear-generated proto-gen gql-gen
