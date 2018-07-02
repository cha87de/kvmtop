package printers

import (
	"log"
	"sort"
	"strings"

	"github.com/cha87de/goncurses"
	"github.com/cha87de/kvmtop/models"
)

var screen *goncurses.Window

// NcursesPrinter describes the ncurses printer
type NcursesPrinter struct {
	models.Printer
}

const HOSTWINHEIGHT = 5
const DOMAINMAXFIELDWIDTH = 8

type KeyValue struct {
	Key   string
	Value string
}

var domainColumnWidths []int

// Open opens the printer
func (printer *NcursesPrinter) Open() {
	// Init goncurses
	var err error
	screen, err = goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	goncurses.Echo(false) // turn echoing of typed characters off
	goncurses.Cursor(0)   // hide cursor
}

// Screen prints the measurements on the screen
func (printer *NcursesPrinter) Screen(printable models.Printable) {

	sortByColumn := 3 // todo: get from input
	maxy, maxx := screen.MaxYX()

	// define host panel
	hostWin, _ := goncurses.NewWindow(HOSTWINHEIGHT, maxx, 0, 0)
	// hostWin.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	goncurses.UpdatePanels()
	goncurses.Update()
	goncurses.NewPanel(hostWin)
	printHost(hostWin, printable.HostFields, printable.HostValues)

	// define domain panel
	domainWin, _ := goncurses.NewWindow(maxy-HOSTWINHEIGHT, maxx, HOSTWINHEIGHT, 0)
	// domainWin.Box(goncurses.ACS_VLINE, goncurses.ACS_HLINE)
	goncurses.UpdatePanels()
	goncurses.Update()
	goncurses.NewPanel(domainWin)
	printDomain(domainWin, printable.DomainFields, printable.DomainValues, sortByColumn)

}

// Close terminates the printer
func (printer *NcursesPrinter) Close() {
	goncurses.End()
}

// CreateNcurses creates a new ncurses printer
func CreateNcurses() NcursesPrinter {
	return NcursesPrinter{}
}

func printHost(window *goncurses.Window, fields []string, values []string) {
	window.Move(1, 1)
	// window.Printf("Whatever: %s", values[0])
	for i, field := range fields {
		window.Printf("%s: %s, ", field, values[i])
	}
}

func printDomain(window *goncurses.Window, fields []string, values map[string][]string, sortByColumn int) {
	window.Move(1, 1)
	// write headline
	currentGroup := ""
	window.Printf("%s", currentGroup)
	for columnID, field := range fields {
		_, cursorPosX := window.CursorYX()

		// split group from field name
		fieldParts := strings.Split(field, "_")
		if len(fieldParts) > 1 {
			groupLabel := strings.Join(fieldParts[0:len(fieldParts)-1], " ")
			if groupLabel != currentGroup {
				// print net group label
				currentGroup = groupLabel
				window.Move(1, cursorPosX)
				window.Printf("%s ", groupLabel)
			}
		}

		// prepare field label length
		fieldLabel := prepareForCell(fieldParts[len(fieldParts)-1], columnID)

		// field label bold if sorted by
		if sortByColumn == columnID {
			window.AttrOn(goncurses.A_BOLD)
		} else {
			window.AttrOff(goncurses.A_BOLD)
		}

		window.AttrOn(goncurses.A_REVERSE)
		window.Move(2, cursorPosX)
		window.Printf("%s ", fieldLabel)
		window.AttrOff(goncurses.A_REVERSE)

	}
	window.Println()

	// create ordered domain list
	domainList := sortDomainIDsByField(values, sortByColumn)

	// iterate over domains
	rowCounter := 3
	for _, domain := range domainList {
		// write domain row
		window.Move(rowCounter, 1)
		for columnID, value := range values[domain.Key] {
			if sortByColumn == columnID {
				window.AttrOn(goncurses.A_BOLD)
			} else {
				window.AttrOff(goncurses.A_BOLD)
			}
			value = prepareForCell(value, columnID)
			window.Printf("%s ", value)
		}
		window.Println()
		rowCounter++
	}
}

func sortDomainIDsByField(values map[string][]string, sortByColumn int) []KeyValue {
	var sorted []KeyValue
	for key, value := range values {
		sorted = append(sorted, KeyValue{key, value[sortByColumn]})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})
	return sorted
}

func prepareForCell(content string, columnID int) string {
	return expandCell(fitInCell(content), columnID)
}

func fitInCell(content string) string {
	if len(content) > DOMAINMAXFIELDWIDTH {
		tmp := strings.Split(content, "")
		content = strings.Join(tmp[0:DOMAINMAXFIELDWIDTH], "")
	}
	return content
}

func expandCell(content string, columnID int) string {
	if len(domainColumnWidths) <= columnID {
		// set current width as column width
		domainColumnWidths = append(domainColumnWidths, len(content))
	} else if len(content) < domainColumnWidths[columnID] {
		// column width is larger, expand with spaces
		spaces := []string{""}
		diff := domainColumnWidths[columnID] - len(content)
		for i := 0; i < diff; i++ {
			spaces = append(spaces, " ")
		}
		content = content + strings.Join(spaces, "")
	} else if len(content) > domainColumnWidths[columnID] {
		// column width is smaller, store larger width
		domainColumnWidths[columnID] = len(content)
	}
	return content
}
