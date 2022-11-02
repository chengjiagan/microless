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
# services that have http interface
REST_SERVICES = \
	socialgraph \
	user \
	composepost \
	usertimeline \
	hometimeline

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

# prefix of image name
DOCKER_PREFIX = waterlemongan/microless-socialnetwork-
# name for build image
BUILD_IMAGE = $(DOCKER_PREFIX)build
# tag for image
VERSION = 0.1.0

docker-build: docker-build-gateway docker-build-test $(foreach s, $(SERVICES), docker-build-$(s))

docker-build-build: docker/Dockerfile.build
	docker build \
		--build-arg SERVICES="$(SERVICES)" \
		-t $(BUILD_IMAGE) -f $< src

docker-build-gateway: docker/Dockerfile.gateway
	docker build \
		-t $(DOCKER_PREFIX)gateway -f $< src

docker-build-test: test/Dockerfile
	docker build \
		-t $(DOCKER_PREFIX)test -f $< test

docker-build-%: docker/Dockerfile.service docker-build-build
	docker build \
		--build-arg BUILD_IMAGE=$(BUILD_IMAGE) \
		--build-arg SERVICE=$* \
		-t $(DOCKER_PREFIX)$* -f $< .

docker-tag: docker-tag-gateway docker-tag-test $(foreach s, $(SERVICES), docker-tag-$(s))

docker-tag-%:
	docker tag $(DOCKER_PREFIX)$* $(DOCKER_PREFIX)$*:$(VERSION)

docker-push: docker-push-gateway docker-push-test $(foreach s, $(SERVICES), docker-push-$(s))

docker-push-%:
	docker push $(DOCKER_PREFIX)$*:$(VERSION)

.PHONY init:
	cd src && go mod download && cd ..
	cd test && pipenv install --dev --system && cd ..
