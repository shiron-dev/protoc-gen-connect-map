package main

import (
	"bytes"
	"fmt"
	"text/template"
	"unicode"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/shiron-dev/protoc-gen-connect-map/gen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, protoFile := range gen.Files {
			if !protoFile.Generate {
				continue
			}
			generate(gen, protoFile)
		}
		return nil
	})
}

func generate(gen *protogen.Plugin, protoFile *protogen.File) {
	filename := protoFile.GeneratedFilenamePrefix + "_connect-map.pb.go"
	g := gen.NewGeneratedFile(filename, protoFile.GoImportPath)

	g.P("package ", protoFile.GoPackageName)
	g.P("")

	for _, service := range protoFile.Services {
		generateAclMap(g, protoFile, service)
	}
}
func getPathName(protoFile *protogen.File, method *protogen.Method) string {
	return fmt.Sprintf("/%s.%s/%s", protoFile.Desc.FullName(), method.Parent.Desc.Name(), method.Desc.Name())
}

func generateAclMap(g *protogen.GeneratedFile, protoFile *protogen.File, service *protogen.Service) bool {
	tmpl := `var {{ .ServiceName }}Map = map[string][]string {
{{- range .Methods }}
	"{{ .Path }}": {
		{{range .Keys }}"{{ . }}",
		{{end }}
	},
{{- end }}
}`

	type MethodData struct {
		Path string
		Keys []string
	}
	type TemplateData struct {
		ServiceName string
		Methods     []MethodData
	}

	var methods []MethodData
	for _, method := range service.Methods {
		path := getPathName(protoFile, method)
		opts := method.Desc.Options().(*descriptorpb.MethodOptions)
		ext := proto.GetExtension(opts, gen.E_ConnectMap)

		optKeys := ext.(*gen.MapOptions).Key
		if len(optKeys) == 0 {
			continue
		}

		methods = append(methods, MethodData{
			Path: path,
			Keys: optKeys,
		})
	}

	tmplData := TemplateData{
		ServiceName: lowerFirstChar(string(service.Desc.Name())),
		Methods:     methods,
	}

	var buf bytes.Buffer
	t := template.Must(template.New("aclMap").Parse(tmpl))
	err := t.Execute(&buf, tmplData)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}

	g.P(buf.String())
	g.P("")

	return true
}

func lowerFirstChar(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
