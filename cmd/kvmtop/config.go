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
	hasCollector := false
	if config.Options.EnableCPU {
		enableCPU()
		hasCollector = true
	}
	if config.Options.EnableMEM {
		enableMEM()
		hasCollector = true
	}
	if config.Options.EnableDISK {
		enableDISK()
		hasCollector = true
	}
	if config.Options.EnableNET {
		enableNET()
		hasCollector = true
	}
	if config.Options.EnableIO {
		enableIO()
		hasCollector = true
	}
	if config.Options.EnableHost {
		enableHOST()
		hasCollector = true
	}

	if !hasCollector {
		// no collector selected, going to add default collector
		enableCPU()
		enableMEM()
	}

	// select printer, ncurse as default.
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

// EnableCPU adds more cpu collector
func enableCPU() {
	collector := cpucollector.CreateCollector()
	models.Collection.Collectors.Store("cpu", &collector)
}

// enableMEM adds more mem collector
func enableMEM() {
	collector := memcollector.CreateCollector()
	models.Collection.Collectors.Store("mem", &collector)
}

// enableDISK adds more disk collector
func enableDISK() {
	collector := diskcollector.CreateCollector()
	models.Collection.Collectors.Store("disk", &collector)
}

// enableNET adds more net collector
func enableNET() {
	collector := netcollector.CreateCollector()
	models.Collection.Collectors.Store("net", &collector)
}

// enableIO adds more io collector
func enableIO() {
	collector := iocollector.CreateCollector()
	models.Collection.Collectors.Store("io", &collector)
}

// enableHOST adds more host collector
func enableHOST() {
	collector := hostcollector.CreateCollector()
	models.Collection.Collectors.Store("host", &collector)
}
