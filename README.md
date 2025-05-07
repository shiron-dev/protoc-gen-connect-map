# Protocol Buffers path mapping for connect-go

This plugin allows you to create a map between the rpc path of connect-go and the option key specified in the protocol buffers

This can be used to specify, for example, protocol buffers-driven interceptors for handling access to specific rpcs.

## Example

```protobuf
syntax = "proto3";

package example;

import "google/protobuf/descriptor.proto";
import "protoc-gen-connect-map/proto/connect_map.proto";

option go_package = "github.com/shiron-dev/protoc-gen-connect-map/example/gen/example;example";

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

```go
package example

var exampleServiceMap = map[string][]string{
	"/example.ExampleService/ExampleMethod": {
		"example_key",
	},
}
```

## How to use

1. Add protoc-gen-connect-map proto to your project

   Add dependency to your buf.yaml

   ```buf.yaml
   deps:
   - buf.build/shiron-dev/protoc-gen-connect-map
   ```

   And locked the dependency

   ```shell
   buf dep update
   ```

2. Install protoc-gen-connect-map

   ```shell
   go install github.com/shiron-dev/protoc-gen-connect-map@latest
   ```

   Add plugin and add override settings to your buf.gen.yaml

   > [!NOTE] 
   > Do not output protoc-gen-connect-map to the same destination as the others
   > Other plugins may respect the overridden go_package_prefix, and paths may be covered.

   ```buf.gen.yaml
   version: v2
   managed:
      enabled: true
      override:
         - file_option: go_package_prefix
            value: github.com/shiron-dev/protoc-gen-connect-map/example/gen/example;example

   plugins:
      - remote: buf.build/protocolbuffers/go
         out: gen/example
         opt: paths=source_relative
         include_imports: true
      - remote: buf.build/connectrpc/go
         out: gen/example
         opt: paths=source_relative
      - local: protoc-gen-connect-map
         out: gen/connect_map
         opt: paths=source_relative

   ```

   - Don't forget to override the managed go_package_prefix!
   - Include_imports in buf.build/protocolbuffers/go must be true
   - It is recommended to set paths=source_relative to simplify the configuration

3. Generate code

   ```shell
   buf generate
   ```
