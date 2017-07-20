package main

import (
	"os"

	"fmt"

	"github.com/cha87de/kvmtop/collectors"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	flags "github.com/jessevdk/go-flags"
)

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&config.Options, flags.Default)
	parser.ShortDescription = "kvmtop"
	parser.LongDescription = "Monitor virtual machine experience from outside on KVM hypervisor level"

	// Parse parameters
	if _, err := parser.Parse(); err != nil {
		fmt.Printf("Error parsing flags: %s", err)
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// Set collectors from flags
	if models.Collection.Collectors == nil {
		models.Collection.Collectors = make(map[string]models.Collector)
	}
	if config.Options.EnableCPU {
		collector := collectors.CreateCollectorCPU()
		models.Collection.Collectors["cpu"] = &collector
	}
	if config.Options.EnableMEM {
		collector := collectors.CreateCollectorMEM()
		models.Collection.Collectors["mem"] = &collector
	}
	if config.Options.EnableDISK {
		collector := collectors.CreateCollectorDISK()
		models.Collection.Collectors["disk"] = &collector
	}
	if config.Options.EnableNET {
		collector := collectors.CreateCollectorNET()
		models.Collection.Collectors["net"] = &collector
	}

}
