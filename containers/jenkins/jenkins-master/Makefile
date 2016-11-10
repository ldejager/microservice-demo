ORGANISATION = rwxlabs
VERSION ?= latest
COMPONENT = jenkins-master

.PHONY: build push release

build:
	docker build --no-cache -t $(ORGANISATION)/$(COMPONENT):$(VERSION) .

push:
	docker push $(ORGANISATION)/$(COMPONENT):$(VERSION)

release: build
	make push -e VERSION=$(VERSION)

default: build
