package main

import (
	"os"

	"fmt"

	"github.com/cha87de/kvmtop/collectors/cpucollector"
	"github.com/cha87de/kvmtop/collectors/diskcollector"
	"github.com/cha87de/kvmtop/collectors/hostcollector"
	"github.com/cha87de/kvmtop/collectors/iocollector"
	"github.com/cha87de/kvmtop/collectors/memcollector"
	"github.com/cha87de/kvmtop/collectors/netcollector"
	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/printers"
	flags "github.com/jessevdk/go-flags"
)

func initializeFlags() {
	// initialize parser for flags
	parser := flags.NewParser(&config.Options, flags.Default)
	parser.ShortDescription = "kvmtop"
	parser.LongDescription = "Monitor virtual machine experience from outside on KVM hypervisor level"

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

	// Set collectors from flags
	if config.Options.EnableCPU {
		collector := cpucollector.CreateCollector()
		models.Collection.Collectors.Store("cpu", &collector)
	}
	if config.Options.EnableMEM {
		collector := memcollector.CreateCollector()
		models.Collection.Collectors.Store("mem", &collector)
	}
	if config.Options.EnableDISK {
		collector := diskcollector.CreateCollector()
		models.Collection.Collectors.Store("disk", &collector)
	}
	if config.Options.EnableNET {
		collector := netcollector.CreateCollector()
		models.Collection.Collectors.Store("net", &collector)
	}
	if config.Options.EnableIO {
		collector := iocollector.CreateCollector()
		models.Collection.Collectors.Store("io", &collector)
	}
	if config.Options.EnableHost {
		collector := hostcollector.CreateCollector()
		models.Collection.Collectors.Store("host", &collector)
	}

	// select printer, ncurse as default.
	if config.Options.PrintBatch { // DEPRECATED remove PrintBatch in future
		printer := printers.CreateText()
		models.Collection.Printer = &printer
	} else {
		switch config.Options.Printer {
		case "ncurses":
			printer := printers.CreateNcurses()
			models.Collection.Printer = &printer
		case "text":
			printer := printers.CreateText()
			models.Collection.Printer = &printer
		case "json":
			printer := printers.CreateJSON()
			models.Collection.Printer = &printer
		default:
			fmt.Println("unknown printer")
			os.Exit(1)
		}

	}

}
