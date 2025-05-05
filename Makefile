MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: init install gen cmd/build.pd.go protoc-gen-connect-map build example

init:
	go mod tidy

install:
	go install ./main.go

cmd/connect_map.pd.go: gen
	buf generate proto

protoc-gen-connect-map: cmd/connect_map.pd.go
	go build -o protoc-gen-connect-map ./main.go

build: protoc-gen-connect-map
	chmod a+x protoc-gen-connect-map

example: build
	cd example && protoc -I. --plugin=$(MAKEFILE_DIR)protoc-gen-connect-map --connect-map_out=. --go_out=. example.proto
