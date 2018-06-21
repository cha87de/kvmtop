package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cha87de/kvmtop/config"
)

// ProcDiskstats defines the fields of one row (one block device) of a /proc/diskstats file
// cf. https://www.kernel.org/doc/Documentation/ABI/testing/procfs-diskstats
type ProcDiskstats struct {

	// 1 - major number
	Majornumber int

	// 2 - minor mumber
	Minornumber int

	// 3 - device name
	Devicename string

	// 4 - reads completed successfully
	Reads uint64

	// 5 - reads merged
	ReadsMerged uint64

	// 6 - sectors read
	SectorsRead uint64

	// 7 - time spent reading (ms)
	TimeReading uint64

	// 8 - writes completed
	Writes uint64

	// 9 - writes merged
	WritesMerged uint64

	//10 - sectors written
	SectorsWritten uint64

	//11 - time spent writing (ms)
	TimeWriting uint64

	//12 - I/Os currently in progress
	CurrentOps uint64

	//13 - time spent doing I/Os (ms)
	TimeForOps uint64

	//14 - weighted time spent doing I/Os (ms)
	WeightedTimeForOps uint64
}

// GetProcDiskstats reads and returns the diskstats for a block device from the proc fs
func GetProcDiskstats(device string) ProcDiskstats {
	stats := ProcDiskstats{Devicename: device}
	filepath := fmt.Sprint(config.Options.ProcFS, "/diskstats")
	filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read proc diskstats: %s\n", err)
		return ProcDiskstats{}
	}

	fmt.Fprintf(os.Stderr, string(filecontent))
	// TODO

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse proc diskstats: %s\n", err)
		return ProcDiskstats{}
	}

	return stats
}
