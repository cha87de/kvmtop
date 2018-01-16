package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

func initializePrinter(wg *sync.WaitGroup) {
	// open configured printer
	models.Collection.Printer.Open()

	// start continuously printing values
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		handleRun()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}

	// close configured printer
	models.Collection.Printer.Close()

	// return from runner
	wg.Done()
}

func handleRun() {
	var fields []string
	var values [][]string

	// collect fields for each collector
	fields = append(fields, "UUID", "name")
	for _, collector := range models.Collection.Collectors {
		output := collector.PrintFields()
		fields = append(fields, output[0:]...)
	}

	// collect values for each domain
	for _, domain := range models.Collection.Domains {
		var domvalues []string
		domvalues = append(domvalues, domain.UUID, domain.Name)
		for _, collector := range models.Collection.Collectors {
			output := collector.PrintValues(domain)
			domvalues = append(domvalues, output[0:]...)
		}
		values = append(values, domvalues)
	}

	models.Collection.Printer.Screen(fields, values)
}
