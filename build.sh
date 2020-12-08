#!/bin/bash

cd `dirname $0`
GOOS=js GOARCH=wasm go build -o server/lib.wasm .