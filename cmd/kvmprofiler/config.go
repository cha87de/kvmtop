package main

import (
	"os"

	"fmt"

	"github.com/cha87de/kvmtop/config"
	flags "github.com/jessevdk/go-flags"
)

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&config.Options, flags.Default)
	parser.ShortDescription = "kvmprofiler"
	parser.LongDescription = "Compute statistical profiles from monitoring data of virtual machines via kvmtop"

	// Parse parameters
	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		if code != 0 {
			fmt.Printf("Error parsing flags: %s", err)
		}
		os.Exit(code)
	}

}
