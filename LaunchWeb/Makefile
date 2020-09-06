TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.0.1-alpha-0.1
BINARY_NAME=LaunchWeb

BINARY_NAME_SERVER=LaunchWeb


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




	


check-env:
	@mkdir -p $(BUILD_FOLDER)/dist/linux/bin
	@mkdir -p $(BUILD_FOLDER)/dist/windows/bin
	@mkdir -p $(BUILD_FOLDER)/dist/arm/bin
	@mkdir -p $(BUILD_FOLDER)/dist/osx/bin
	cp -R config $(BUILD_FOLDER)/dist/linux/
	cp -R config $(BUILD_FOLDER)/dist/windows/
	cp -R config $(BUILD_FOLDER)/dist/arm/
	cp -R config $(BUILD_FOLDER)/dist/osx/
	cp -R extras $(BUILD_FOLDER)/dist/linux/
	


## Linting
lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...


## Make linux packages
package-linux:
	cd build/dist/ && tar zcvf linux-dist.tar.gz linux/
	cd build/dist/ && zip -9 linux-dist.zip -r linux/


versioning:
	./version.sh ${VERSION} ${TIME}



##Compile Web Resources 
compile/webresources: versioning
	go-bindata-assetfs  -nometadata -pkg=ui -ignore=bindata.go webAssets/...
	mv bindata.go pkg/web/ui/bindata.go


build/web-build:
	cd webAssets && npm install && npm run build

build/LaunchWeb-linux:
	cd cmd/RunWebApp && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/LaunchWeb .

build/development: compile/webresources build/LaunchWeb-linux
	build/dist/linux/bin/LaunchWeb --config cfg/sample.cfg


distribute: 
	./release.sh ${VERSION} ${TIME}


