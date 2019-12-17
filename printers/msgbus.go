package printers

import (
	"fmt"
	"kvmtop/models"
	"os"
	"strings"
	"time"

	drModels "github.com/disresc/lib/models"
	transmitter "github.com/disresc/lib/transmitter"
)

// CreateMsgbus creates a new simple text printer
func CreateMsgbus() MsgbusPrinter {
	return MsgbusPrinter{}
}

// MsgbusPrinter prints to a Go-Micro message bus
type MsgbusPrinter struct {
	models.Printer
	transmitterService *transmitter.Service
}

// Open opens the printer
// called once, must return!
func (printer *MsgbusPrinter) Open() {
	hostname, _ := os.Hostname()
	printer.transmitterService = transmitter.NewService(fmt.Sprintf("kvmtop-%s", hostname))
	printer.transmitterService.RegisterRequestCallback(func(request *drModels.Request) bool {
		if request.GetSource() == "hosts" || request.GetSource() == fmt.Sprintf("host-%s", hostname) || request.GetSource() == "ves" {
			// source match!
			transmitterParts := strings.Split(request.GetTransmitter(), "-")
			if len(transmitterParts) != 2 {
				// transmitter format not matching
				return false
			}
			if transmitterParts[0] != "kvmtop" {
				// transmitter not kvmtop
				return false
			}
			collectors := []string{"cpu", "mem", "net", "disk", "io", "host", "psi"}
			foundCollector := false
			for _, v := range collectors {
				if v == transmitterParts[1] {
					foundCollector = true
					break
				}
			}
			if !foundCollector {
				// collector not found
				return false
			}

			// all checks valid so far!
			// accept the request
			return true
		}
		return false
	})
	printer.transmitterService.Start()
	// go printer.listen()
}

// Screen prints the measurements on the screen
// this is called every config.Options.Frequency seconds
func (printer *MsgbusPrinter) Screen(printable models.Printable) {
	// transform printables into event

	printer.sendHostEvents(printable)
	printer.sendVEEvents(printable)

}

// Close terminates the printer
func (printer *MsgbusPrinter) Close() {
	printer.transmitterService.Close()
}

func (printer *MsgbusPrinter) listen() {
	// TODO implement push instead of periodic pull

	//for update := range models.Collection.Updates {
	//fmt.Printf("going to publish new event %v\n", update.Event)
	//printer.transmitterService.Publish(update.Event)
	//}
}

func (printer *MsgbusPrinter) sendHostEvents(printable models.Printable) {
	itemsByTransmitter := make(map[string][]*drModels.EventItem)
	for i, field := range printable.HostFields {
		metricParts := strings.Split(field, "_")
		collector := metricParts[0]
		transmitter := fmt.Sprintf("kvmtop-%s", collector)

		if _, exists := itemsByTransmitter[transmitter]; !exists {
			itemsByTransmitter[transmitter] = make([]*drModels.EventItem, 0)
		}

		item := drModels.EventItem{
			Transmitter: transmitter,
			Metric:      field,
			Value:       printable.HostValues[i],
		}

		itemsByTransmitter[transmitter] = append(itemsByTransmitter[transmitter], &item)
	}
	hostname, _ := os.Hostname()
	source := fmt.Sprintf("host-%s", hostname)
	interval := 10

	for _, txItems := range itemsByTransmitter {
		event := drModels.Event{
			Source:    source,
			Timestamp: time.Now().Unix(),
			Items:     txItems,
		}
		printer.transmitterService.Publish(&event, interval)
	}
}

func (printer *MsgbusPrinter) sendVEEvents(printable models.Printable) {
	for domID, domvalue := range printable.DomainValues {
		itemsByTransmitter := make(map[string][]*drModels.EventItem)
		for i, value := range domvalue {
			field := printable.DomainFields[i]

			metricParts := strings.Split(field, "_")
			collector := metricParts[0]
			transmitter := fmt.Sprintf("kvmtop-%s", collector)

			item := drModels.EventItem{
				Transmitter: transmitter,
				Metric:      field,
				Value:       value,
			}
			itemsByTransmitter[transmitter] = append(itemsByTransmitter[transmitter], &item)
		}
		source := fmt.Sprintf("ve-%s", domID)
		interval := 10

		for _, txItems := range itemsByTransmitter {
			event := drModels.Event{
				Source:    source,
				Timestamp: time.Now().Unix(),
				Items:     txItems,
			}
			printer.transmitterService.Publish(&event, interval)
		}
	}

}
