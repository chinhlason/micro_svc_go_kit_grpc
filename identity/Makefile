DSN="postgresql://postgres:postgres@localhost:5432/db1?sslmode=disable"
PROTO_FILE=identity.proto

g-up:
	goose -dir ./migrations postgres $(DSN) up

g-down:
	goose -dir ./migrations postgres $(DSN) down

p-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/identity.proto

g-gen:
	@protoc -I . --proto_path=protobuf $(PROTO_FILE) \
        --go_out=./protobuf --go_opt paths=source_relative \
        --go-grpc_out=./protobuf --go-grpc_opt paths=source_relative \
        --grpc-gateway_out=./protobuf --grpc-gateway_opt paths=source_relative

apg-gen:
	@protoc -I . --proto_path=protobuf $(PROTO_FILE) \
		--go_out=../gateway/protobuf --go_opt paths=source_relative \
		--go-grpc_out=../gateway/protobuf --go-grpc_opt paths=source_relative \
		--grpc-gateway_out=../gateway/protobuf --grpc-gateway_opt paths=source_relative