#!/bin/bash

dstArch="$GOARCH"
if [ "$dstArch" = "386" ]; then
    dstArch="i386"
fi

dstName="file-encoding-converter"
if [ "$GOOS" = "windows" ]; then
    dstName="$dstName.exe"
fi

echo "Building for $GOOS $dstArch"
go build -o dist/$GOOS/$dstArch/$dstName ./src
