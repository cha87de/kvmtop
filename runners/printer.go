package runners

import (
	"sync"
	"time"

	"github.com/cha87de/kvmtop/config"
	"github.com/cha87de/kvmtop/models"
)

var collectors []string

func initializePrinter(wg *sync.WaitGroup) {
	// open configured printer
	models.Collection.Printer.Open()

	// define collectors and their order
	models.Collection.Collectors.Map.Range(func(key, collectorRaw interface{}) bool {
		collectorName := key.(string)
		collectors = append(collectors, collectorName)
		return true
	})

	// start continuously printing values
	start := time.Now()
	for n := 0; config.Options.Runs == -1 || n < config.Options.Runs; n++ {
		// sleep before execution
		nextRun := start.Add(time.Duration(config.Options.Frequency) * time.Second)
		time.Sleep(nextRun.Sub(time.Now()))
		handleRun()
		start = time.Now()
	}

	// close configured printer
	models.Collection.Printer.Close()

	// return from runner
	wg.Done()
}

func handleRun() {
	printable := models.Printable{}

	// add general domain fields first
	printable.DomainFields = []string{"UUID", "name"}
	printable.DomainValues = make(map[string][]string)
	models.Collection.Domains.Map.Range(func(key, value interface{}) bool {
		uuid := key.(string)
		domain := value.(models.Domain)
		printable.DomainValues[uuid] = []string{
			uuid,
			domain.Name,
		}
		return true
	})

	// collect fields for each collector and merge together
	for _, collectorName := range collectors { // BUG: concurrent map iteration and map write
		collector, ok := models.Collection.Collectors.Load(collectorName)
		if !ok {
			continue
		}
		collectorPrintable := collector.Print()

		// merge host data
		printable.HostFields = append(printable.HostFields, collectorPrintable.HostFields[0:]...)
		printable.HostValues = append(printable.HostValues, collectorPrintable.HostValues[0:]...)

		// merge domain data
		printable.DomainFields = append(printable.DomainFields, collectorPrintable.DomainFields[0:]...)
		for uuid := range collectorPrintable.DomainValues {
			printable.DomainValues[uuid] = append(printable.DomainValues[uuid], collectorPrintable.DomainValues[uuid][0:]...)
		}
	}

	models.Collection.Printer.Screen(printable)
}
