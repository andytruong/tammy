install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go get -d entgo.io/ent/cmd/ent

install.1:
	go install entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc

gen.1:
	# === Generate ent SDKs ===
	go run entgo.io/ent/cmd/ent generate ./schema

gen.2:
	# === Generate proto files ===
	go run cmd/gen/main.go

gen.3:
	# === Generate GRPC implementation ===
	go generate tammy/ent/proto/entpb

gen: gen.1 gen.2 gen.3
	go fmt ./...
