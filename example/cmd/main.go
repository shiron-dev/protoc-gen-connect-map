package main

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	connect_map "github.com/shiron-dev/protoc-gen-connect-map/example/gen/connect_map"
	example "github.com/shiron-dev/protoc-gen-connect-map/example/gen/example"
	"github.com/shiron-dev/protoc-gen-connect-map/example/gen/example/exampleconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type ExampleServer struct{}

func (s *ExampleServer) ExampleMethod(
	ctx context.Context,
	req *connect.Request[example.ExampleRequest],
) (*connect.Response[example.ExampleResponse], error) {
	res := connect.NewResponse(&example.ExampleResponse{
		Message: connect_map.ExampleServiceMap["/example.ExampleService/ExampleMethod"][0],
	})
	return res, nil
}

func (s *ExampleServer) ExampleMethod2(
	ctx context.Context,
	req *connect.Request[example.ExampleRequest],
) (*connect.Response[example.ExampleResponse], error) {
	return s.ExampleMethod(ctx, req)
}

func (s *ExampleServer) ExampleMethod3(
	ctx context.Context,
	req *connect.Request[example.ExampleRequest],
) (*connect.Response[example.ExampleResponse], error) {
	return s.ExampleMethod(ctx, req)
}

func main() {
	server := &ExampleServer{}
	mux := http.NewServeMux()
	path, handler := exampleconnect.NewExampleServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
