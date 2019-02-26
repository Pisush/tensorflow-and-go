
WDIR := /go/src/github.com/Pisush/tensorflow-and-go
DIR := ${CURDIR}:${WDIR}

CONTAINER_NAME := tensorflow-and-go
DOCKER_IMAGE := dahernan/tensorflow-and-go

login:
	docker run -i -v $(DIR) -w $(WDIR) --entrypoint=/bin/bash --name $(CONTAINER_NAME) -t $(DOCKER_IMAGE)

dockerbuild:
	docker build -f Dockerfile -t $(DOCKER_IMAGE) .

down:
	docker rm -f $(CONTAINER_NAME) 2>/dev/null || true

dockerpush:
	docker push $(DOCKER_IMAGE):latest

PHONY: dockerbuild login down dockerpush