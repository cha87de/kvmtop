package printers

import (
	"fmt"
	drModels "github.com/disresc/lib/models"
	transmitter "github.com/disresc/lib/transmitter"
	"kvmtop/models"
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
	printer.transmitterService = transmitter.NewService("kvmtop-hostxy")
	printer.transmitterService.Start()
	// go printer.listen()
}

// Screen prints the measurements on the screen
// this is called every config.Options.Frequency seconds
func (printer *MsgbusPrinter) Screen(printable models.Printable) {
	//fmt.Printf("received printables: %v", printable)

	// transform printables into event

	for i, field := range printable.HostFields {
		event := drModels.Event{
			Name:  field,
			Value: printable.HostValues[i],
		}
		fmt.Printf("going to send event: %v\n", event)
		printer.transmitterService.Publish(event)
	}

}

// Close terminates the printer
func (printer *MsgbusPrinter) Close() {
	printer.transmitterService.Close()
}

func (printer *MsgbusPrinter) listen() {
	for update := range models.Collection.Updates {
		fmt.Printf("going to publish new event %v\n", update.Event)
		printer.transmitterService.Publish(update.Event)
	}
}
