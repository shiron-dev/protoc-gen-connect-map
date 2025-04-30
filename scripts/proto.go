package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed schema.proto.tpl
var SchemaStr []byte

type TemplateData struct {
	GoPackage string
}

func main() {
	f, err := os.Create("proto/build.proto")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = f.Chmod(0644)
	if err != nil {
		panic(err)
	}

	data := TemplateData{
		GoPackage: "./gen;gen",
	}
	t := template.Must(template.New("schema.proto.tpl").Parse(string(SchemaStr)))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	f.Close()

	f, err = os.Create("proto/schema.proto")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = f.Chmod(0644)
	if err != nil {
		panic(err)
	}

	data = TemplateData{
		GoPackage: "github.com/shiron-dev/protoc-gen-connect-map/gen",
	}
	t = template.Must(template.New("schema.proto.tpl").Parse(string(SchemaStr)))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	f.Close()
}
