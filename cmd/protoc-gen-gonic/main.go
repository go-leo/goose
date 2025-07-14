package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-leo/gonic/cmd/protoc-gen-gonic/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

var flags flag.FlagSet

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), "v1.6.1")
		os.Exit(0)
	}
	options := &protogen.Options{ParamFunc: flags.Set}
	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate(plugin)
	})
}

func generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}
		if len(file.Services) <= 0 {
			continue
		}
		gen, err := generator.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := gen.Generate(); err != nil {
			return err
		}
	}
	return nil
}
