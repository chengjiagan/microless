# prefix of image name
DOCKER_NAME = waterlemongan/microless-loadbalancer-base

docker-build: docker/Dockerfile
	docker build \
		-t $(DOCKER_NAME) -f $< src
