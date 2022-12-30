#! /bin/bash

name="mqtt-publisher"
suffix="-linux-amd64"
file_name="${name}${suffix}"

echo "Removing old files"
rm -rf "../${file_name}"

echo "build start."
env GOOS=linux GOARCH=amd64 go build -v .././main.go
mv main "../${file_name}"
echo "build succeeded."