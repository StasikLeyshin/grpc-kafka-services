goCmd ?= go

DIR_ROOT_PROTO=api

DIR_MANAGE_SERVER_SERVICE=server_v1

gen_grpc_proto = protoc --proto_path=$(DIR_ROOT_PROTO)/$(1) \
                         --go_out=pkg/$(1) \
                         --go_opt=paths=source_relative \
                         --go-grpc_out=pkg/$(1) \
                         --go-grpc_opt=paths=source_relative \
                         $(DIR_ROOT_PROTO)/$(1)/*.proto

all: gen_grpc_manage_server_service


gen_grpc_manage_server_service:
	$(call gen_grpc_proto,$(DIR_MANAGE_SERVER_SERVICE))
