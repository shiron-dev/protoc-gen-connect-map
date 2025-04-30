MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: gen cmd/build.pd.go protoc-gen-connect-map build example

gen:
	go run ./scripts/proto.go

cmd/build.pd.go: gen
	protoc -I. --go_out=. proto/build.proto

protoc-gen-connect-map: cmd/build.pd.go
	go build -o protoc-gen-connect-map ./cmd/main.go

build: protoc-gen-connect-map
	chmod a+x protoc-gen-connect-map

example: build
	cd example && protoc -I. --plugin=$(MAKEFILE_DIR)protoc-gen-connect-map --connect-map_out=. --go_out=. example.proto
