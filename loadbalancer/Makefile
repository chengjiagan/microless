# directory where protobuf definitions stay
PROTO_DIR = proto

proto: proto-loadbalancer

proto-loadbalancer: $(PROTO_DIR)/loadbalancer.proto
	protoc --proto_path=. --proto_path=proto \
		--go-grpc_out=src --go-grpc_opt=module="microless/loadbalancer" \
		--go_out=src --go_opt=module="microless/loadbalancer" \
		$<

# prefix of image name
DOCKER_NAME = waterlemongan/microless-loadbalancer-base
# tag for image
VERSION = 0.1.0

docker-build: docker/Dockerfile
	docker build \
		-t $(DOCKER_NAME) -f $< src

docker-tag:
	docker tag $(DOCKER_NAME) $(DOCKER_NAME):$(VERSION)

docker-push:
	docker push $(DOCKER_NAME):$(VERSION)
