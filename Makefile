install.gen-protobuf:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

install.gen-grpc: install.gen-protobuf
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

mod.vendor: install.gen-grpc
	go fmt ./...
	go mod tidy
	go mod vendor

api.gen:
	protoc --proto_path=./pkg/proto/ \
		--go_out=./pkg/proto/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/proto/pb \
		--go-grpc_opt=paths=source_relative \
		./pkg/proto/rpc.proto

# ---------------------
# Services › Store
# ---------------------
# ent init --target=pkg/ent/schema User

ent.gen-sdk:
	rm -rf ./pkg/store/*
	ent generate --target=./pkg/store ./ent
	go run cmd/gen/grpc.go
	# protoc --proto_path=./pkg/proto/ --go_out=./pkg/pb --go_opt=paths=source_relative --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative ./pkg/proto/*.proto
