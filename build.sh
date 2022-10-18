#!/bin/bash

echo "Removing old builds"
rm -f bin/*

echo "Building binaries for various architectures"
for GOOS in darwin linux windows; do
    for GOARCH in 386 amd64; do
      if [ "$GOOS" == 'darwin' ] && [ "$GOARCH" == '386' ]; then
        echo "Skipping $GOOS $GOARCH"
      else
        go build -v -ldflags="-s -w" -o bin/pcc-reporter-$GOOS-$GOARCH
      fi
    done
done

echo "Rename windows executables"
for x in ./bin/pcc-reporter-windows-*; do mv $x $x.exe; done