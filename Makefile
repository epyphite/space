TIME=$(shell date +"%Y%m%d.%H%M%S")
VERSION=0.2.alpha
BINARY_NAME=nasaExplorer

BINARY_NAME_SERVER=nasaExplorer


BASE_FOLDER = $(shell pwd)
BUILD_FOLDER  = $(shell pwd)/build


FLAGS_LINUX   = CGO_LDFLAGS="-L./LIB -Wl,-rpath -Wl,\$ORIGIN/LIB" CGO_ENABLED=1 GOOS=linux GOARCH=amd64  
FLAGS_DARWIN  = OSXCROSS_NO_INCLUDE_PATH_WARNINGS=1 MACOSX_DEPLOYMENT_TARGET=10.6 CC=o64-clang CXX=o64-clang++ CGO_ENABLED=0
FLAGS_FREEBSD = GOOS=freebsd GOARCH=amd64 CGO_ENABLED=1
FLAGS_WINDOWS = GOOS=windows GOARCH=amd64 CC=i686-w64-mingw32-gcc CGO_ENABLED=1 
FLAGS_ARM = CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 CC=arm-linux-gnueabi-gcc

GOFLAGS_WINDOWS = -ldflags -H=windowsgui


clean:
	rm -rvf build/dist/

## Make linux packages
package-linux:
	cd build/dist/ && tar zcvf linux-dist.tar.gz linux/
	cd build/dist/ && zip -9 linux-dist.zip -r linux/


## Linux Build

build/nasaExplorer-linux:
	cd NASA/cmd/ && ${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/nasaExplorer .


build/LaunchAPI-linux: 
	${FLAGS_LINUX} go build -o ${BUILD_FOLDER}/dist/linux/bin/LaunchAPI LaunchAPI/cmd/main.go
	cp -R LaunchAPI/data ${BUILD_FOLDER}/dist/linux/bin/



distribute:
	./upload_github.sh ${VERSION} ${TIME}


build/launchapi-container:
	docker build -t launchapi -f Dockerfile-LaunchApi .

build/nasaapi-container:
	docker build -t nasaapi -f Dockerfile-NasaApi .

build/launchweb-container:
	docker build -t lalaunchwebunchapi -f Dockerfile-LaunchWeb .

clean:
	rm -Rvf build/dist/

gen_proto:
	./proto-gen.sh ${BUILD_FOLDER}
