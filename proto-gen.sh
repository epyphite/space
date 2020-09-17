
BASE_OUTPUT=$1

GOLANG_OUTPUT="$BASE_OUTPUT/golang/pkg/api/v1/"
PYTHON_OUTPUT="$BASE_OUTPUT/python/pkg/api/v1/"
CSHARP_OUTPUT="$BASE_OUTPUT/csharp/pkg/api/v1/"
CPLUS_OUTPUT="$BASE_OUTPUT/cplus/pkg/api/v1/"
echo "Generating Directories"

mkdir -p $GOLANG_OUTPUT
mkdir -p $PYTHON_OUTPUT
mkdir -p $CSHARP_OUTPUT
mkdir -p $CPLUS_OUTPUT

echo "Generating orbit GO Structure"
protoc --proto_path=api/space/orbital/launch/v1 \
        --proto_path=api/space/orbital/launch/v1 \
        --go_out=plugins=grpc:$GOLANG_OUTPUT \
        --go_opt=paths=source_relative \
            orbit.proto
echo "Generating User Python Structure"
protoc --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1 --python_out=$PYTHON_OUTPUT orbit.proto
echo "Generating User C# Structure"
protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --csharp_out=$CSHARP_OUTPUT orbit.proto
echo "Generating User C++ Structure"
protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --cpp_out=$CPLUS_OUTPUT orbit.proto


echo "Generating orbit GO Structure"
protoc --proto_path=api/space/orbital/launch/v1 \
        --proto_path=api/space/orbital/launch/v1 \
        --go_out=plugins=grpc:$GOLANG_OUTPUT \
        --go_opt=paths=source_relative \
        spaceport.proto
echo "Generating User Python Structure"
protoc --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1 --python_out=$PYTHON_OUTPUT spaceport.proto
echo "Generating User C# Structure"
protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --csharp_out=$CSHARP_OUTPUT spaceport.proto
echo "Generating User C++ Structure"
protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --cpp_out=$CPLUS_OUTPUT spaceport.proto




echo "Generating service GO Structure"
#protoc --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1 --go_out=plugins=grpc:$GOLANG_OUTPUT launchService.proto
protoc --proto_path=api/space/orbital/launch/v1 \
       --proto_path=api/space/orbital/launch/v1 \
       --go_out=plugins=grpc:$GOLANG_OUTPUT \
        launchService.proto

protoc --proto_path=api/space/orbital/launch/v1  \
        --proto_path=api/space/orbital/launch/v1 \
        --grpc-gateway_out=logtostderr=true:$GOLANG_OUTPUT \
        launchService.proto


echo "Generating Launch Service Python Structure"
#protoc --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1 --grpc_python_out=$PYTHON_OUTPUT --python_out=$PYTHON_OUTPUT launchService.proto

python -m grpc_tools.protoc -Iapi/space/orbital/launch/v1 --python_out=$PYTHON_OUTPUT --grpc_python_out=$PYTHON_OUTPUT launchService.proto


echo "Generating  Launch Service C# Structure"
protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --csharp_out=$CSHARP_OUTPUT launchService.proto

echo "Generating  Launch Service C++ Structure"
#protoc  --proto_path=api/space/orbital/launch/v1 --proto_path=api/space/orbital/launch/v1  --grpc_out=$CPLUS_OUTPUT --cpp_out=$CPLUS_OUTPUT launchService.proto


protoc -I api/space/orbital/launch/v1 --grpc_out=$CPLUS_OUTPUT --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` launchService.proto
protoc -I api/space/orbital/launch/v1 --cpp_out=$CPLUS_OUTPUT launchService.proto