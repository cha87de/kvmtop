package printers

import (
	"fmt"
	"strconv"

	"github.com/cha87de/kvmtop/models"
)

// JSONPrinter describes the json printer
type JSONPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *JSONPrinter) Open() {
	OutputOpen()
}

// Screen prints the measurements on the screen
func (printer *JSONPrinter) Screen(printable models.Printable) {
	Output(fmt.Sprintf("{ \"host\": {"))
	hostFields := printable.HostFields
	hostValues := printable.HostValues
	i := 0
	for _, value := range hostValues {
		if i > 0 {
			Output(fmt.Sprintf(","))
		}

		// but """ only for strings
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			Output(fmt.Sprintf("\"%s\": %d", hostFields[i], intValue))
		} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			Output(fmt.Sprintf("\"%s\": %f", hostFields[i], floatValue))
		} else {
			Output(fmt.Sprintf("\"%s\": \"%s\"", hostFields[i], value))
		}
		i++
	}

	Output(fmt.Sprintf("}, \"domains\": ["))

	domainFields := printable.DomainFields
	domainValues := printable.DomainValues
	i = 0
	for domvalue := range domainValues {
		if i > 0 {
			Output(fmt.Sprintf(","))
		}
		Output(fmt.Sprintf("{"))
		for j, value := range domainValues[domvalue] {
			if j > 0 {
				Output(fmt.Sprintf(","))
			}

			// but """ only for strings
			if _, err := strconv.ParseInt(value, 10, 64); err == nil {
				Output(fmt.Sprintf("\"%s\": %s", domainFields[j], value))
			} else if _, err := strconv.ParseFloat(value, 64); err == nil {
				Output(fmt.Sprintf("\"%s\": %s", domainFields[j], value))
			} else {
				Output(fmt.Sprintf("\"%s\": \"%s\"", domainFields[j], value))
			}
		}
		Output(fmt.Sprintf("}"))
		i++
	}
	Output(fmt.Sprintf("]}\n"))
}

// Close terminates the printer
func (printer *JSONPrinter) Close() {
	OutputClose()
}

// CreateJSON creates a new simple text printer
func CreateJSON() JSONPrinter {
	return JSONPrinter{}
}
