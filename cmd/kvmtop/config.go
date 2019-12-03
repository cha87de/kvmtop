package main

import (
	"os"

	"fmt"

	flags "github.com/jessevdk/go-flags"
	"kvmtop/collectors/cpucollector"
	"kvmtop/collectors/diskcollector"
	"kvmtop/collectors/hostcollector"
	"kvmtop/collectors/iocollector"
	"kvmtop/collectors/memcollector"
	"kvmtop/collectors/netcollector"
	"kvmtop/collectors/psicollector"
	"kvmtop/config"
	"kvmtop/models"
	"kvmtop/printers"
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
	if config.Options.EnablePressure {
		enablePressure()
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
	case "msgbus":
		printer := printers.CreateMsgbus()
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

// enablePressure adds more pressure collector
func enablePressure() {
	collector := psicollector.CreateCollector()
	models.Collection.Collectors.Store("pressure", &collector)
}

// enableHOST adds more host collector
func enableHOST() {
	collector := hostcollector.CreateCollector()
	models.Collection.Collectors.Store("host", &collector)
}
