package main

import (
	"os"

	"github.com/cha87de/kvmtop/config"
	flags "github.com/jessevdk/go-flags"
)

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&config.Options, flags.Default)
	parser.ShortDescription = "kvmtop"
	parser.LongDescription = "Monitor virtual machine experience from outside on KVM hypervisor level"

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

}
