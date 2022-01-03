#!/bin/bash

APP_NAME=saber
VERSION=1.0.0

fmt:
	goimports -l -w .

clean:
	rm -rf output/
  
build:clean fmt
	go build -o output/${APP_NAME} .

setup:clean fmt
	go build -ldflags "-s -w" -o output/${APP_NAME} .

doc:
	go run . doc

pkg:clean linux window mac tar

linux:fmt
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o output/linux/${APP_NAME} .

window:fmt
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-s -w' -o output/window/${APP_NAME}.exe .

mac:fmt
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-s -w' -o output/mac/${APP_NAME} .

macM1:fmt
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -ldflags '-s -w' -o output/mac/${APP_NAME} .


tar:
	tar -zcvf output/${APP_NAME}$(VERSION).window-amd64.tar.gz output/window/${APP_NAME}.exe
	tar -zcvf output/${APP_NAME}$(VERSION).linux-amd64.tar.gz output/linux/${APP_NAME}
	tar -zcvf output/${APP_NAME}$(VERSION).darwin-amd64.tar.gz output/mac/${APP_NAME}
