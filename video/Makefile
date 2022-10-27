# all functions
SERVICES = \
	video \
	preview \
	split \
	transcode \
	merge
# functions that do not require extra packages
NORMAL_SERVICES = \
	video
# functions that require ffmpeg
FFMPEG_SERVICES = \
	preview \
	split \
	transcode \
	merge

# directory where dockerfiles stay
DOCKER_DIR = docker
# prefix of image name
DOCKER_PREFIX = microless/video-
# name for build image
BUILD_IMAGE = $(DOCKER_PREFIX)build
# tag for image
VERSION = 0.1.0

requirements: $(foreach s, $(SERVICES), requirements-$(s))

requirements-test:
	cd test && pipenv lock && pipenv requirements > requirements.txt

requirements-%:
	cd src/$* && pipenv lock && pipenv requirements > requirements.txt

docker-build-test:
	docker build -t $(DOCKER_PREFIX)test test

docker-build: $(foreach s, $(NORMAL_SERVICES), docker-build-normal-$(s)) $(foreach s, $(FFMPEG_SERVICES), docker-build-ffmpeg-$(s))

docker-build-ffmpeg-%: docker/Dockerfile.ffmpeg
	docker build \
		--build-arg SERVICE=$* \
		-t $(DOCKER_PREFIX)$* -f $< src/$*

docker-build-normal-%: docker/Dockerfile
	docker build \
		--build-arg SERVICE=$* \
		-t $(DOCKER_PREFIX)$* -f $< src/$*

docker-tag: $(foreach s, $(SERVICES), docker-tag-$(s))

docker-tag-%:
	docker tag $(DOCKER_PREFIX)$* $(DOCKER_PREFIX)$*:$(VERSION)
