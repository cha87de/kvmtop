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
	for collectorName := range models.Collection.Collectors {
		collectors = append(collectors, collectorName)
	}

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
	printable := models.Printable{}

	// add general domain fields first
	printable.DomainFields = []string{"UUID", "name"}
	printable.DomainValues = make(map[string][]string)
	for uuid := range models.Collection.Domains {
		printable.DomainValues[uuid] = []string{
			uuid,
			models.Collection.Domains[uuid].Name,
		}
	}

	// collect fields for each collector and merge together
	for _, collectorName := range collectors {
		collector := models.Collection.Collectors[collectorName]
		collectorPrintable := collector.Print(models.Collection.Host, models.Collection.Domains)

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
