package main

import (
	"flag"
	"log"
	"strings"

	"github.com/mwantia/nautilus-debug/pkg/core"
	"github.com/mwantia/nautilus/pkg/plugin"
)

var (
	Address = flag.String("address", "", "Define the address used to serve this plugin over network.")
)

func main() {
	flag.Parse()

	impl := core.NewImpl()

	if strings.TrimSpace(*Address) != "" {
		if err := plugin.RegisterPipelineProcessor(impl, *Address); err != nil {
			log.Fatalf("Unable to register plugin: %v", err)
		}
	} else {
		if err := plugin.ServePipelineProcessor(impl); err != nil {
			log.Fatalf("Unable to serve plugin: %v", err)
		}
	}
}
