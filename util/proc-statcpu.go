package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cha87de/kvmtop/config"
)

// ProcStatCPU describes one CPU row in /proc/stat
type ProcStatCPU struct {
	Name string

	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	IRQ       uint64
	SoftIRQ   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

// GetProcStatCPU reads and returns the cpu related rows in /proc/stat
func GetProcStatCPU() []ProcStatCPU {
	stats := []ProcStatCPU{}
	filepath := fmt.Sprint(config.Options.ProcFS, "/stat")

	file, err := os.Open(filepath)
	if err != nil {
		// cannot open file ...
		return stats
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	format := "%s %d %d %d %d %d %d %d %d %d %d"

	for scanner.Scan() {
		row := scanner.Text()
		stat := ProcStatCPU{}

		// filter rows, only consider cpu rows
		if !strings.HasPrefix(row, "cpu") {
			continue
		}

		_, err := fmt.Sscanf(
			string(row), format,
			&stat.Name,
			&stat.User,
			&stat.Nice,
			&stat.System,
			&stat.Idle,
			&stat.IOWait,
			&stat.IRQ,
			&stat.SoftIRQ,
			&stat.Steal,
			&stat.Guest,
			&stat.GuestNice,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse row in proc stat: %s\n", err)
			continue
		}
		stats = append(stats, stat)
	}
	return stats
}
