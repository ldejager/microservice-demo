ORGANISATION = rwxlabs
VERSION ?= latest
COMPONENT = jenkins-volumes

.PHONY: build push release

pull:
	docker pull jenkinsci/jenkins

build:
	docker build --no-cache -t $(ORGANISATION)/$(COMPONENT):$(VERSION) .

push:
	docker push $(ORGANISATION)/$(COMPONENT):$(VERSION)

release: pull build
	make push -e VERSION=$(VERSION)

default: build
