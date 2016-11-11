ORGANISATION = ldejager
VERSION ?= latest
COMPONENT = microservices-demo

.PHONY: build push release

build:
		docker build --no-cache -t $(ORGANISATION)/$(COMPONENT):$(VERSION) .

push:
		docker push $(ORGANISATION)/$(COMPONENT):$(VERSION)

release: build
		make push -e VERSION=$(VERSION)

default: build
