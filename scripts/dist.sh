#!/bin/bash

SCRIPTS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PROJECT_DIR="$SCRIPTS_DIR/.."
BUILD_DIR="$PROJECT_DIR/build"
cd "$PROJECT_DIR/cmd/app"
env GOOS=linux GOARCH=386 go build -o "$PROJECT_DIR/build/linux-386/app"
env GOOS=linux GOARCH=amd64 go build -o "$PROJECT_DIR/build/linux-amd64/app"
env GOOS=windows GOARCH=386 go build -o "$PROJECT_DIR/build/windows-386/app.exe"
env GOOS=windows GOARCH=amd64 go build -o "$PROJECT_DIR/build/windows-amd64/app.exe"
env GOOS=darwin GOARCH=amd64 go build -o "$PROJECT_DIR/build/darwin-amd64/app"
cd $PROJECT_DIR