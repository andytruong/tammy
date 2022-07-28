mod.vendor:
	go fmt ./...
	go mod tidy
	go mod vendor
