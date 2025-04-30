# Protocol Buffers path mapping for connect-go

This plugin allows you to create a map between the rpc path of connect-go and the option key specified in the protocol buffers

This can be used to specify, for example, protocol buffers-driven interceptors for handling access to specific rpcs.

## Example

```protobuf
syntax = "proto3";

package example;

import "google/protobuf/descriptor.proto";
import "protoc-gen-connect-map/proto/connect_map.proto";

option go_package = "example.protobuf;example";

message ExampleRequest {
  string name = 1;
}
message ExampleResponse {
  string message = 1;
}

service ExampleService {
  rpc ExampleMethod(ExampleRequest) returns (ExampleResponse) {
    option (connect_map.connect_map) = {
      key: "example_key"
    };
  }
}

```

```golang
package example

var exampleServiceMap = map[string][]string{
	"/example.ExampleService/ExampleMethod": {
		"example_key",
	},
}
```

## How to use

1. Add schema.proto to your project(Recommended to use git submodule)

   ```shell
   git submodule add https://github.com/shiron-dev/protoc-gen-connect-map.git protoc-gen-connect-map
   ```

1. Install protoc-gen-connect-map or build it yourself

   1. Install the plugin using go install

      ```shell
      go install github.com/shiron-dev/protoc-gen-connect-map@latest
      ```

   1. Build the plugin yourself(use git submodule)

      ```shell
      cd protoc-gen-connect-map
      make build
      ```

1. Import schema.proto in your proto file

   ```protobuf
   import "protoc-gen-connect-map/proto/connect_map.proto";
   ```

1. Generate the code using protoc

   1. Install the plugin using go install

      ```shell
      protoc --connect-map_out=. your.proto
      ```


   1. Build the plugin yourself(use git submodule)

      ```shell
      protoc --plugin=protoc-gen-connect-map/protoc-gen-connect-map --connect-map_out=. your.proto
      ```
