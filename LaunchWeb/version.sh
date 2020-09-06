#!/usr/bin/env bash
TIME=$2
VERSION=$1

echo "$TIME"
sed -i "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${VERSION}\"/" pkg/constants/version.go
sed -i "s/BuildTime    string = \"[^\"]*\"/BuildTime    string = \"${TIME}\"/" pkg/constants/version.go

sed -i "s/BuildVersion string = \"[^\"]*\"/BuildVersion string = \"${VERSION}\"/" pkg/constants/version.go
sed -i "s/BuildTime    string = \"[^\"]*\"/BuildTime    string = \"${TIME}\"/" pkg/constants/version.go