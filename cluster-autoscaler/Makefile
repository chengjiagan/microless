# prefix of image name
DOCKER_NAME = waterlemongan/microless-cluster-autoscaler
# tag for image
VERSION = 0.1.0

docker-build: docker/Dockerfile
	docker build \
		-t $(DOCKER_NAME) -f $< src

docker-tag:
	docker tag $(DOCKER_NAME) $(DOCKER_NAME):$(VERSION)

docker-push:
	docker push $(DOCKER_NAME):$(VERSION)
