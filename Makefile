.PHONY: gen-proto

# https://github.com/aurora-is-near/devops-stuff/tree/main/docker/protoc
gen-proto:
	docker run --rm -v $(PWD)/:/proto protoc \
		--proto_path=/proto/ \
		--go_out=/proto/. \
		--go-vtproto_out=/proto/. \
		messagebackup/messagebackup.proto
