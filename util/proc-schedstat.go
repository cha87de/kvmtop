package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/cha87de/kvmtop/config"
)

// ProcSchedStat defines the fields of a /proc/[pid]/schedstat file
// cf. https://www.kernel.org/doc/Documentation/scheduler/sched-stats.txt
type ProcSchedStat struct {
	// The process ID.
	PID int
	// time spent on the cpu
	Cputime uint64
	// time spent waiting on a runqueue
	Runqueue uint64
	// # of timeslices run on this cpu
	Timeslices uint64
}

// GetProcSchedStat reads and returns the schedstat for a process from the proc fs
func GetProcSchedStat(pid int) ProcSchedStat {
	stats := ProcSchedStat{PID: pid}
	filepath := fmt.Sprint(config.Options.ProcFS, "/", strconv.Itoa(pid), "/schedstat")
	filecontent, _ := ioutil.ReadFile(filepath)

	_, err := fmt.Fscan(
		bytes.NewBuffer(filecontent),
		&stats.Cputime,
		&stats.Runqueue,
		&stats.Timeslices,
	)

	if err != nil {
		return ProcSchedStat{}
	}

	return stats
}
