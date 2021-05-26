
TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.0.1-alpha-0.2
BINARY_NAME=satComOne

BINARY_NAME_SERVER=satComOne


BASE_FOLDER = $(shell pwd)
BUILD_FOLDER  = $(shell pwd)/build

FLAGS_LINUX   = CGO_LDFLAGS="-L./LIB -Wl,-rpath -Wl,\$ORIGIN/LIB" CGO_ENABLED=1 GOOS=linux GOARCH=amd64  
FLAGS_DARWIN  = OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.6 CC=o64-clang CXX=o64-clang++ CGO_ENABLED=0
FLAGS_FREEBSD = GOOS=freebsd GOARCH=amd64 CGO_ENABLED=1
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 
FLAGS_ARM = CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 CC=arm-linux-gnueabi-gcc

GOFLAGS_WINDOWS = -ldflags -H=windowsgui


init: 
	@mkdir -vp $(BASE_FOLDER)/build
	@mkdir -vp $(BASE_FOLDER)/cmd
	@mkdir -vp $(BASE_FOLDER)/config
	@mkdir -vp $(BASE_FOLDER)/extras
	@mkdir -vp $(BASE_FOLDER)/pkg/constants
	@mkdir -vp $(BASE_FOLDER)/third-party
	@mkdir -vp $(BASE_FOLDER)/api
	@mkdir -vp $(BASE_FOLDER)/vendor
	@echo "Creating Base Files" 
	@touch $(BASE_FOLDER)/pkg/constants/version.go
	@touch $(BASE_FOLDER)/vendor/packages_windows.txt
	@touch $(BASE_FOLDER)/vendor/packages_linux.txt



genproto:
	protoc --proto_path=./pkg/icd --go_out=./pkg/api/v1/ ./pkg/icd/radio.proto
	protoc --proto_path=./pkg/icd --go_out=./pkg/api/v1/ ./pkg/icd/log.proto
	protoc --proto_path=./pkg/icd --go_out=./pkg/api/v1/ ./pkg/icd/ping.proto
	protoc --proto_path=./pkg/icd --go_out=./pkg/api/v1/ ./pkg/icd/status.proto


build/satcom-linux:
	cd cmd/ && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/satComOne .


