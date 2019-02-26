
WDIR := /go/src/github.com/Pisush/tensorflow-and-go
DIR := ${CURDIR}:${WDIR}

DOCKER_IMAGE := dahernan/tensorflow-and-go

login:
	docker run -i -v $(DIR) -w $(WDIR) --entrypoint=/bin/bash -t $(DOCKER_IMAGE)

dockerbuild:
	docker build -f Dockerfile -t $(DOCKER_IMAGE) .

dockerpush:
	docker push $(DOCKER_IMAGE):latest

PHONY: dockerbuild login dockerpush