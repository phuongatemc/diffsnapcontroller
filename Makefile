IMAGE_NAME ?= quay.io/isim/diffsnap-controller
IMAGE_TAG ?= latest
DOCKER_BUILDKIT ?= 1

build:
	DOCKER_BUILDKIT=$(DOCKER_BUILDKIT) docker build --rm -t $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
