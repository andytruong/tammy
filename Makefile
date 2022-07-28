mod.vendor:
	go fmt ./...
	go mod tidy
	go mod vendor

module.ent.start:
	# === start ent server ===
	go run tammy/ent/cmd/server

dev: module.ent.start
