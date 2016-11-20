ORGANISATION = ldejager
VERSION ?= latest
COMPONENT = microservices-demo

.PHONY: build

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

all: deps build

deps:
	go get -u github.com/ldejager/microservice-demo

build: build_static build_cross build_tar build_sha

build_static:
	go install github.com/ldejager/microservice-demo
	mkdir -p release
	cp $(GOPATH)/bin/microservice-demo release/

build_cross:
	GOOS=linux   GOARCH=amd64 CGO_ENABLED=0 go build -o release/linux/amd64/microservice-demo   github.com/ldejager/microservice-demo
	GOOS=darwin  GOARCH=amd64 CGO_ENABLED=0 go build -o release/darwin/amd64/microservice-demo  github.com/ldejager/microservice-demo

build_tar:
	tar -cvzf release/linux/amd64/microservice-demo.tar.gz   -C release/linux/amd64   microservice-demo
	tar -cvzf release/darwin/amd64/microservice-demo.tar.gz  -C release/darwin/amd64  microservice-demo

build_sha:
	sha256sum release/linux/amd64/microservice-demo.tar.gz   > release/linux/amd64/microservice-demo.sha256
	sha256sum release/darwin/amd64/microservice-demo.tar.gz  > release/darwin/amd64/microservice-demo.sha256

default: all
