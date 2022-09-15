# all microservices
SERVICES = \
	poststorage \
	usertimeline \
	user \
	socialgraph \
	hometimeline \
	media \
	urlshorten \
	usermention \
	text \
	composepost

# directory where protobuf definitions stay
PROTO_DIR = proto
# protobuf definition for data models
PROTO_DATA = $(PROTO_DIR)/data.proto

proto-data: $(PROTO_DATA)
	protoc --proto_path=. --proto_path=proto \
		--go_out=src \
		--go_opt=module="microless/socialnetwork" \
		$<

proto-%: $(PROTO_DIR)/%.proto
	protoc --proto_path=. --proto_path=proto \
		--go-grpc_out=src --go-grpc_opt=module="microless/socialnetwork" \
		--go_out=src --go_opt=module="microless/socialnetwork" \
		$<

proto-rest-%: $(PROTO_DIR)/%.proto
	protoc --proto_path=. --proto_path=proto \
		--grpc-gateway_out src \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt module="microless/socialnetwork" \
		$<

proto-test-data: $(PROTO_DATA)
	protoc --proto_path=. --proto_path=proto \
		--pyi_out=test \
		$<
	python3 -m grpc_tools.protoc \
		--proto_path=. --proto_path=proto \
		--python_out=test \
		$<

proto-test-%: $(PROTO_DIR)/%.proto
	protoc --proto_path=. --proto_path=proto \
		--pyi_out=test \
		$<
	python3 -m grpc_tools.protoc \
		--proto_path=. --proto_path=proto \
		--python_out=test \
		--grpc_python_out=test \
		$<

proto: proto-data $(foreach s, $(SERVICES), proto-$(s)) $(foreach s, $(REST_SERVICES), proto-rest-$(s))

proto-test: proto-test-data $(foreach s, $(SERVICES), proto-test-$(s))
	touch test/proto/__init__.py

# directory where dockerfiles stay
DOCKER_DIR = docker
# name for build image
BUILD_IMAGE = microless/build
# tag for image
VERSION = 0.1.3

docker: $(foreach s, $(SERVICES), docker-$(s)) $(foreach s, $(REST_SERVICES), docker-rest-$(s)) docker-gateway

docker-build: docker/Dockerfile.build
	docker build \
		--build-arg SERVICES="$(SERVICES)" \
		-t $(BUILD_IMAGE):$(VERSION) -f $< src

docker-gateway: docker/Dockerfile.gateway src/gateway/
	docker build \
		-t microless/gateway:$(VERSION) -f $< src

docker-%: docker/Dockerfile.service docker-build src/%/
	docker build \
		--build-arg BUILD_IMAGE=$(BUILD_IMAGE):$(VERSION) \
		--build-arg SERVICE=$* \
		-t microless/$*:$(VERSION) -f $< .

.PHONY init:
	cd src && go mod download && cd ..
	cd test && pipenv install --dev --system && cd ..
