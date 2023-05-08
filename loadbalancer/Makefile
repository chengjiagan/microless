# prefix of image name
DOCKER_NAME = waterlemongan/microless-loadbalancer-base
# tag for image
VERSION = 0.1.0

docker-build: docker/Dockerfile
	docker build \
		-t $(DOCKER_NAME) -f $< src
