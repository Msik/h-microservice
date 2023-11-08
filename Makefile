PHONY: .vendor-proto
.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api
	cp api/api.proto vendor.protogen/api
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
		mkdir -p vendor.protogen/github.com/envoyproxy &&\
		git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate ;\
	fi

PHONY: generate
generate:
	protoc -I vendor.protogen \
			--go_out=pkg/api --go_opt=paths=import \
			--go-grpc_out=pkg/api --go-grpc_opt=paths=import \
			--grpc-gateway_out=pkg/api \
			--grpc-gateway_opt=logtostderr=true \
			--grpc-gateway_opt=paths=import \
			--validate_out lang=go:pkg/api \
			--swagger_out=allow_merge=true,merge_file_name=api:swagger \
			api/api.proto

PHONY: cpProto
cpProto:
	cp api/api.proto vendor.protogen/api
