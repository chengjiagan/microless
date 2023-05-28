# all microservices
SERVICES = \
	ping \
	pong
# services that have http interface
REST_SERVICES = \
	ping

# directory where protobuf definitions stay
PROTO_DIR = proto

proto-%: $(PROTO_DIR)/%.proto
	protoc --proto_path=. --proto_path=proto --proto_path=/usr/include \
		--go-grpc_out=src --go-grpc_opt=module="microless/pingpong" \
		--go_out=src --go_opt=module="microless/pingpong" \
		$<

proto-rest-%: $(PROTO_DIR)/%.proto
	protoc --proto_path=. --proto_path=proto --proto_path=/usr/include \
		--grpc-gateway_out src \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt module="microless/pingpong" \
		$<

proto: $(foreach s, $(SERVICES), proto-$(s)) $(foreach s, $(REST_SERVICES), proto-rest-$(s))

# prefix of image name
DOCKER_PREFIX = waterlemongan/microless-pingpong-
# name for base build image
BASE_IMAGE = waterlemongan/microless-loadbalancer-base
# name for build image
BUILD_IMAGE = $(DOCKER_PREFIX)build
# tag for image
VERSION = 0.1.0

docker-build: docker-build-gateway $(foreach s, $(SERVICES), docker-build-$(s))

docker-build-build: docker/Dockerfile.build
	docker build \
		--build-arg BASE_IMAGE="$(BASE_IMAGE)" \
		--build-arg SERVICES="gateway $(SERVICES)" \
		-t $(BUILD_IMAGE) -f $< src

docker-build-%: docker/Dockerfile.service docker-build-build
	docker build \
		--build-arg BUILD_IMAGE=$(BUILD_IMAGE) \
		--build-arg SERVICE=$* \
		-t $(DOCKER_PREFIX)$* -f $< .

docker-tag: docker-tag-gateway $(foreach s, $(SERVICES), docker-tag-$(s))

docker-tag-%:
	docker tag $(DOCKER_PREFIX)$* $(DOCKER_PREFIX)$*:$(VERSION)

docker-push: docker-push-gateway $(foreach s, $(SERVICES), docker-push-$(s))

docker-push-%:
	docker push $(DOCKER_PREFIX)$*:$(VERSION)

docker-minikube: docker-minikube-gateway $(foreach s, $(SERVICES), docker-minikube-$(s))

docker-minikube-%:
	minikube image load $(DOCKER_PREFIX)$*:$(VERSION)
