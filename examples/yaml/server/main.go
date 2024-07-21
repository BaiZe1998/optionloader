package main

import (
	"fmt"
	"github.com/kitex-contrib/optionloader/optionloader/server"
	"github.com/kitex-contrib/optionloader/source/yaml"
)

func main() {
	filePath := "./examples/yaml/server/config.yaml"
	cfg, err := yaml.LoadServerConfig(filePath)
	if err != nil {
		panic(err)
	}
	loader := server.NewServerOptionLoader(cfg)
	options, err := loader.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(options))
}
