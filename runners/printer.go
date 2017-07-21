package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
	"github.com/cha87de/kvmtop/printers"
)

func initializePrinter(wg *sync.WaitGroup) {
	// print header
	// TODO

	// start continuously printing values
	for n := -1; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		start := time.Now()
		print()
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
	}
	wg.Done()
}

func print() {
	for _, domain := range models.Collection.Domains {
		var values []string
		values = append(values, domain.UUID, domain.Name)
		for _, collector := range models.Collection.Collectors {
			output := collector.Print(domain)
			values = append(values, output[0:]...)
		}
		printers.Textprint(values)
	}
}
