TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.1-alpha
BINARY_NAME=esaExplorer

BINARY_NAME_SERVER=esaExplorer


BASE_FOLDER = $(shell pwd)
BUILD_FOLDER  = $(shell pwd)/../build


all: build/docker
test: lint build/docker/unit-test




FLAGS_LINUX   = CGO_LDFLAGS="-L./LIB -Wl,-rpath " CGO_ENABLED=1 GOOS=linux GOARCH=amd64  
FLAGS_DARWIN  = OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.6 CC=o64-clang CXX=o64-clang++ CGO_ENABLED=0
FLAGS_FREEBSD = GOOS=freebsd GOARCH=amd64 CGO_ENABLED=1
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 
FLAGS_ARM = CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 CC=arm-linux-gnueabi-gcc

GOFLAGS_WINDOWS = -ldflags -H=windowsgui

getdeps:
	./getDeps.sh


versioning:
	./version.sh ${VERSION} ${TIME}



## Linting
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...

## Linux Build

build/esaExplorer-linux: 
	${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/esaExplorer cmd/main.go


package:
	#cp cfg/sample.cfg $(BUILD_FOLDER)/dist/linux/
	cd build/dist/ && tar zcvf linux-dist.tar.gz linux/
	cd build/dist/ && zip -9 linux-dist.zip -r linux/


build/dev: build/esaExplorer-linux package

build/esaExplorer-container:
	docker build -t esaExplorer -f Dockerfile.api .
	
distribute: 
	./release.sh ${VERSION} ${TIME}



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

clean:
	rm -Rvf build/dist/
