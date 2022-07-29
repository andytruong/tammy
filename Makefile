gen.client.wip:
	mkdir -p resources/proto/
	curl -sS https://gitlab.com/vn/tammy/ent/-/raw/v0.02/proto/entpb/entpb.proto --output resources/proto/ent.proto
	protoc something
