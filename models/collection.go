package models

import (
	drModels "github.com/disresc/lib/models"
)

func init() {
	// initialize the collection variable
	Collection.Domains = *NewDomains()
	Collection.Collectors = *NewCollectors()
	Collection.Host = Host{
		Measurable: NewMeasurable(),
	}
	Collection.Updates = make(chan Update)
}

// Collection of domains and other stuff
var Collection struct {
	Host           Host
	Domains        Domains
	Collectors     Collectors
	Printer        Printer
	LibvirtDomains LibvirtDomains
	Updates        chan Update
}

// SendUpdate sends an update if receivers are attached to Collection.Updates
// channel
func SendUpdate(update Update) {
	select {
	case Collection.Updates <- update:
		// update sent
	default:
		// not sent since no receiver
	}
}

// SendUpdateEvent takes an Event and sends it to the channel
func SendUpdateEvent(event drModels.Event) {
	SendUpdate(Update{
		Event: event,
	})
}

// Update encapsulates a collection or lookup update event
type Update struct {
	Event drModels.Event
}
