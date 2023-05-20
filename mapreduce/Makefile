# all microservices
SERVICES = \
	mapreduce \
	file \
	mapper \
	reducer

# directory where dockerfiles stay
DOCKER_DIR = docker
# prefix of image name
DOCKER_PREFIX = waterlemongan/microless-mapreduce-
# name for build image
BUILD_IMAGE = $(DOCKER_PREFIX)build
# tag for image
VERSION = 0.1.0

requirements: requirements-test $(foreach s, $(SERVICES), requirements-$(s))

requirements-test:
	cd test && pipenv lock && pipenv requirements > requirements.txt

requirements-%:
	cd src/$* && pipenv lock && pipenv requirements > requirements.txt

docker-build: docker-build-test $(foreach s, $(SERVICES), docker-build-$(s))

docker-build-test: test/Dockerfile
	docker build \
		-t $(DOCKER_PREFIX)test -f $< test

docker-build-%: docker/Dockerfile
	docker build \
		--build-arg SERVICE=$* \
		-t $(DOCKER_PREFIX)$* -f $< src/$*

docker-tag: docker-tag-test $(foreach s, $(SERVICES), docker-tag-$(s))

docker-tag-%:
	docker tag $(DOCKER_PREFIX)$* $(DOCKER_PREFIX)$*:$(VERSION)

docker-push: docker-push-test $(foreach s, $(SERVICES), docker-push-$(s))

docker-push-%:
	docker push $(DOCKER_PREFIX)$*:$(VERSION)
