#!/bin/bash

go mod tidy
echo "Building Project..."
go build -o ./bin/
echo "Build success."
