TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.2.alpha
BINARY_NAME=satExplorer
BINARY_NAME_SERVER=satExplorer
BASE_FOLDER = $(shell pwd)
BUILD_FOLDER  = $(shell pwd)/../build

all: build/docker
test: lint build/docker/unit-test



## Linting
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...


infrastructure:
	docker pull mongo
	docker run -d -p 27017-27019:27017-27019 --name SpaceMongo mongo

## Make linux packages
package-linux:
	cd build/dist/ && tar zcvf linux-dist.tar.gz linux/
	cd build/dist/ && zip -9 linux-dist.zip -r linux/


versioning:
	./version.sh ${VERSION} ${TIME}

## Linux Build

build/satExplorer-linux:
	cd cmd/ && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/satExplorer .

distribute:
	./upload_github.sh ${VERSION} ${TIME}



PLATFORM=local

.PHONY: build/docker
build/docker:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

.PHONY: build/docker/unit-test
build/docker/unit-test:
	@docker build . --target unit-test

.PHONY: build/docker/unit-test-coverage
build/docker/unit-test-coverage:
	@docker build . --target unit-test-coverage \
	--output coverage/
	cat coverage/cover.out