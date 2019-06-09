package util

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cha87de/kvmtop/config"
)

// ProcDiskstat defines the fields of one row (one block device) of a /proc/diskstats file
// cf. https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats
type ProcDiskstat struct {

	// 1 - major number
	Majornumber int

	// 2 - minor mumber
	Minornumber int

	// 3 - device name
	Devicename string

	// 4 - reads completed successfully f1
	Reads uint64

	// 5 - reads merged f2
	ReadsMerged uint64

	// 6 - sectors read f3
	SectorsRead uint64

	// 7 - time spent reading (ms) f4
	TimeReading uint64

	// 8 - writes completed f5
	Writes uint64

	// 9 - writes merged f6
	WritesMerged uint64

	//10 - sectors written f7
	SectorsWritten uint64

	//11 - time spent writing (ms) f8
	TimeWriting uint64

	//12 - I/Os currently in progress f9
	CurrentOps uint64

	//13 - time spent doing I/Os (ms) f10
	TimeForOps uint64

	//14 - weighted time spent doing I/Os (ms) f11
	WeightedTimeForOps uint64
}

// GetProcDiskstats reads and returns the diskstats from the proc fs
func GetProcDiskstats() map[string]ProcDiskstat {
	stats := make(map[string]ProcDiskstat)

	filepath := fmt.Sprint(config.Options.ProcFS, "/diskstats")
	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	format := "%d %d %s %d %d %d %d %d %d %d %d %d %d %d"

	for scanner.Scan() {
		row := scanner.Text()
		diskstat := ProcDiskstat{}

		_, err := fmt.Sscanf(
			string(row), format,
			&diskstat.Majornumber,
			&diskstat.Minornumber,
			&diskstat.Devicename,
			&diskstat.Reads,
			&diskstat.ReadsMerged,
			&diskstat.SectorsRead,
			&diskstat.TimeReading,
			&diskstat.Writes,
			&diskstat.WritesMerged,
			&diskstat.SectorsWritten,
			&diskstat.TimeWriting,
			&diskstat.CurrentOps,
			&diskstat.TimeForOps,
			&diskstat.WeightedTimeForOps,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse row in proc diskstats: %s\n", err)
			continue
		}
		stats[diskstat.Devicename] = diskstat
	}

	return stats
}
