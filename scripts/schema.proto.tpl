syntax = "proto3";

package connect_map;

import "google/protobuf/descriptor.proto";

{{ if .GoPackage -}}
option go_package = "{{ .GoPackage }}";
{{ end -}}

extend google.protobuf.MethodOptions {
  MapOptions connect_map = 50101;
}

message MapOptions {
  repeated string key = 1;
}
