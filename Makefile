IMAGE_NAME ?= quay.io/isim/diffsnap-controller
IMAGE_TAG ?= latest
DOCKER_BUILDKIT ?= 1

compile:
	rm -f ./diffsnap-controller
	CGO_ENABLED=0 GOOS=linux go build -o ./diffsnap-controller .

build:
	DOCKER_BUILDKIT=$(DOCKER_BUILDKIT) docker build --rm -t $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

codegen: vendor
	./hack/update-codegen.sh

compile-aggapi: vendor
	rm -rf diffsnap-aggapi
	CGO_ENABLED=0 GOOS=linux go build -o diffsnap-aggapi ./cmd/aggapi/...

codegen-aggapi: vendor
	./hack/update-codegen-aggapi.sh

vendor:
	go mod vendor
