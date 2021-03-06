package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"kvmtop/collectors/cpucollector"
	"kvmtop/collectors/iocollector"
	"kvmtop/collectors/netcollector"
	"kvmtop/config"
	"kvmtop/models"
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
			fmt.Printf("Error parsing flags: %s\n", err)
		}
		os.Exit(code)
	}

	// Set collectors from flags
	if config.Options.EnableCPU {
		collector := cpucollector.CreateCollector()
		models.Collection.Collectors.Store("cpu", &collector)
	}
	if config.Options.EnableMEM {
		fmt.Println("memory profiling not supported.")
	}
	if config.Options.EnableDISK {
		fmt.Println("disk profiling not supported.")
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
		fmt.Println("host profiling not supported.")
	}

	// Parse periodsize csv string
	periodSizeStr := strings.Split(config.Options.Profiler.PeriodSize, ",")
	periodSize := make([]int, 0)
	for _, s := range periodSizeStr {
		if s == "" {
			continue
		}
		si, _ := strconv.Atoi(s)
		periodSize = append(periodSize, si)
	}
	config.Options.Profiler.PeriodSizeParsed = periodSize

}
