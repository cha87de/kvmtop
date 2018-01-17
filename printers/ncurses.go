package printers

import (
	"bufio"
	"bytes"
	"log"

	tw "text/tabwriter"

	"github.com/cha87de/goncurses"
	"github.com/cha87de/kvmtop/models"
)

var outputHelpers struct {
	screen                *goncurses.Window
	tabwriterBuffer       bytes.Buffer
	tabwriterBufferWriter bufio.Writer
	tabwriter             tw.Writer
}

// NcursesPrinter describes the ncurses printer
type NcursesPrinter struct {
	models.Printer
}

// Open opens the printer
func (printer *NcursesPrinter) Open() {
	// Init goncurses
	s, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	outputHelpers.screen = s
	goncurses.Echo(false) // turn echoing of typed characters off
	goncurses.Cursor(0)   // hide cursor

	// declare buffer and writer
	outputHelpers.tabwriterBuffer = *new(bytes.Buffer)
	outputHelpers.tabwriterBufferWriter = *bufio.NewWriter(&outputHelpers.tabwriterBuffer)
	outputHelpers.tabwriter = *new(tw.Writer)
	outputHelpers.tabwriter.Init(&outputHelpers.tabwriterBufferWriter, 0, 8, 1, ' ', 0)
}

// Screen prints the measurements on the screen
func (printer *NcursesPrinter) Screen(fields []string, values [][]string) {

	var buffer bytes.Buffer
	var rawline string

	// write headline
	for _, field := range fields {
		buffer.WriteString(field)
		buffer.WriteString("\t")
	}
	buffer.WriteString("\n")
	rawline = buffer.String()
	outputHelpers.tabwriter.Write([]byte(rawline))

	// iterate over domains
	for _, domvalue := range values {
		buffer.Reset()
		for _, value := range domvalue {
			buffer.WriteString(value)
			buffer.WriteString("\t")
		}
		buffer.WriteString("\n")
		rawline = buffer.String()
		outputHelpers.tabwriter.Write([]byte(rawline))
	}

	outputHelpers.tabwriter.Flush()
	outputHelpers.tabwriterBufferWriter.Flush()
	tablines := outputHelpers.tabwriterBuffer.String()

	outputHelpers.screen.MovePrint(0, 0, tablines)
	outputHelpers.screen.Refresh()
	outputHelpers.tabwriterBuffer.Reset()
}

// Close terminates the printer
func (printer *NcursesPrinter) Close() {
	goncurses.End()
}

// CreateNcurses creates a new ncurses printer
func CreateNcurses() NcursesPrinter {
	return NcursesPrinter{}
}
