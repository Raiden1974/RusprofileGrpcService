.DEFAULT_GOAL := all

PACKAGES_WITH_TESTS:=$(shell go list -f="{{if or .TestGoFiles .XTestGoFiles}}{{.ImportPath}}{{end}}" ./... | grep -v '/vendor/')
TEST_TARGETS:=$(foreach p,${PACKAGES_WITH_TESTS},test-$(p))

GO:=go

include test.mk

PROTOC_BUILD_DIR:=services_api/build
PROTOC_SOURCE_DIR:=services_api/source
PROTOC_SOURCE_SUBDIR:=$(shell find $(PROTOC_SOURCE_DIR) -maxdepth 1 -type d)

TEST_COMPOSE:=docker-compose.tests.yaml
PROTOC_GEN_TS_PATH:="$(shell npm root -g)/grpc_tools_node_protoc_ts/bin/protoc-gen-ts"

PROTOC_BUILD_DIR_GO:=$(PROTOC_BUILD_DIR)/go
PROTOC_BUILD_DIR_JS:=$(PROTOC_BUILD_DIR)/js
PROTOC_BUILD_DIR_CURL:=$(PROTOC_BUILD_DIR)/grpcurl

.PHONY: all
all: testall

.PHONY: clean
clean:
	rm -rf tmp
	rm -rf database/client/session/mocks
	rm -rf $(PROTOC_BUILD_DIR)

.PHONY: govendor
govendor: clean
	go get github.com/golang/dep/cmd/dep
	dep ensure -v

.PHONY: proto
proto: govendor

	# update code generator
	#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	#go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/gogo/protobuf/proto
	go get -u github.com/gogo/protobuf/protoc-gen-gogofaster
	go get -u github.com/gogo/protobuf/gogoproto


	# Golang
	$(foreach dir, $(PROTOC_SOURCE_SUBDIR), \
		mkdir -m 755 -p $(PROTOC_BUILD_DIR_GO)/$(shell basename $(dir)) && \
		protoc \
		--proto_path /usr/local/include \
		--proto_path $(PROTOC_SOURCE_DIR) \
		--proto_path vendor \
		--proto_path $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path $(dir) \
    	--gogofaster_out=Mgoogle/protobuf/empty.proto=github.com/golang/protobuf/ptypes/empty,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,goproto_registration=true,plugins=grpc:$(PROTOC_BUILD_DIR_GO) \
		$(dir)/*.proto;)

	# this is for using with grpcurl util
	# go get github.com/fullstorydev/grpcurl
	# go install github.com/fullstorydev/grpcurl/cmd/grpcurl
	# then grpcurl -protoset ~/go_home/src/github.com/PROFITVenchurs/scLib2/services_api/sc_cli_api/protoset.bin -d '{}' URL_TO_SERVICE sc_cli_api.CliApi/GetActivePorts
	$(foreach dir, $(PROTOC_SOURCE_SUBDIR), \
		mkdir -m 755 -p $(PROTOC_BUILD_DIR_CURL)/$(shell basename $(dir)) && \
		protoc \
		--proto_path /usr/local/include \
		--proto_path $(PROTOC_SOURCE_DIR) \
		--proto_path vendor \
		--proto_path $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path $(dir) \
		--include_imports \
    	--descriptor_set_out=$(PROTOC_BUILD_DIR_CURL)/$(shell basename $(dir))/protoset.bin \
		$(dir)/*.proto;)

	# Node. Before will be install. :
	# sudo npm install -g google-protobuf
	# sudo npm install -g grpc-tools --unsafe-perm
	# sudo npm install -g grpc_tools_node_protoc_ts --unsafe-perm

	mkdir -m 755 -p $(PROTOC_BUILD_DIR_JS)

	protoc \
	--proto_path /usr/local/include \
	--proto_path $(PROTOC_SOURCE_DIR) \
	--proto_path vendor \
	--proto_path $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	$(PROTOC_SOURCE_DIR)/*/*.proto \
	--js_out=import_style=commonjs,binary:$(PROTOC_BUILD_DIR_JS) \
	--grpc_out=$(PROTOC_BUILD_DIR_JS) \
	--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin`

	protoc \
	--proto_path /usr/local/include \
	--proto_path $(PROTOC_SOURCE_DIR) \
	--proto_path vendor \
	--proto_path $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	$(PROTOC_SOURCE_DIR)/*/*.proto \
	--plugin="protoc-gen-ts=$(PROTOC_GEN_TS_PATH)" \
	--js_out="import_style=commonjs,binary:$(PROTOC_BUILD_DIR_JS)" \
	--ts_out="service=true:$(PROTOC_BUILD_DIR_JS)"

	# fix imports in js files
	# remove nullable annotation for golang
	find  $(PROTOC_BUILD_DIR_JS) -type f -name "*.js" -a -exec sed -i -e 's/var github_com_gogo_protobuf_gogoproto_gogo_pb =/\/\//g' {} +
	find  $(PROTOC_BUILD_DIR_JS) -type f -name "*.ts" -a -exec sed -i -e 's/import \* as github_com_gogo_protobuf_gogoproto_gogo_pb from/\/\//g' {} +
	find  $(PROTOC_BUILD_DIR_JS) -type f -name "*.js" -a -exec sed -i -e 's/goog.object.extend(proto, github_com_gogo_protobuf_gogoproto_gogo_pb);//g' {} +

.PHONY: mocks
mocks: proto
	go get github.com/vektra/mockery/.../
	mockery -name=IDbClient -dir=database/client/session -recursive -output=database/client/session/mocks


.PHONY: testafter
testafter:
	docker-compose -f $(TEST_COMPOSE) stop

.PHONY: testbefore
testbefore: testafter
	docker-compose -f $(TEST_COMPOSE) up

.PHONY: testall
testall: mocks test
