#!/bin/bash

APP_NAME=saber

fmt:
	goimports -l -w .

clean:
	rm -rf bin/
  
build:clean fmt
	go build -o bin/${APP_NAME} .
