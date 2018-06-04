package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

// ProcStat defines the fields of a /proc/[pid]/stat file
// cf. http://man7.org/linux/man-pages/man5/proc.5.html
type ProcStat struct {
	// The process ID.
	PID int
	// The filename of the executable.
	Comm string
	// The process state.
	State string
	// The PID of the parent of this process.
	PPID int
	// The process group ID of the process.
	PGRP int
	// The session ID of the process.
	Session int
	// The controlling terminal of the process.
	TTY int
	// The ID of the foreground process group of the controlling terminal of
	// the process.
	TPGID int
	// The kernel flags word of the process.
	Flags uint
	// The number of minor faults the process has made which have not required
	// loading a memory page from disk.
	MinFlt uint
	// The number of minor faults that the process's waited-for children have
	// made.
	CMinFlt uint
	// The number of major faults the process has made which have required
	// loading a memory page from disk.
	MajFlt uint
	// The number of major faults that the process's waited-for children have
	// made.
	CMajFlt uint
	// Amount of time that this process has been scheduled in user mode,
	// measured in clock ticks.
	UTime uint
	// Amount of time that this process has been scheduled in kernel mode,
	// measured in clock ticks.
	STime uint
	// Amount of time that this process's waited-for children have been
	// scheduled in user mode, measured in clock ticks.
	CUTime uint
	// Amount of time that this process's waited-for children have been
	// scheduled in kernel mode, measured in clock ticks.
	CSTime uint
	// For processes running a real-time scheduling policy, this is the negated
	// scheduling priority, minus one.
	Priority int
	// The nice value, a value in the range 19 (low priority) to -20 (high
	// priority).
	Nice int
	// Number of threads in this process.
	NumThreads int
	// The time the process started after system boot, the value is expressed
	// in clock ticks.
	Starttime uint64
	// Virtual memory size in bytes.
	VSize int
	// Resident set size in pages.
	RSS int
}

// GetProcStat reads and returns the stat for a process from the proc fs
func GetProcStat(pid int) ProcStat {
	stats := ProcStat{PID: pid}
	filepath := fmt.Sprint("/proc/", strconv.Itoa(pid), "/stat")
	filecontent, _ := ioutil.ReadFile(filepath)
	// fmt.Printf("%s", filecontent)

	var (
		ignore int

		l = bytes.Index(filecontent, []byte("("))
		r = bytes.LastIndex(filecontent, []byte(")"))
	)

	if l < 0 || r < 0 {
		return ProcStat{}
	}

	stats.Comm = string(filecontent[l+1 : r])
	_, err := fmt.Fscan(
		bytes.NewBuffer(filecontent[r+2:]),
		&stats.State,
		&stats.PPID,
		&stats.PGRP,
		&stats.Session,
		&stats.TTY,
		&stats.TPGID,
		&stats.Flags,
		&stats.MinFlt,
		&stats.CMinFlt,
		&stats.MajFlt,
		&stats.CMajFlt,
		&stats.UTime,
		&stats.STime,
		&stats.CUTime,
		&stats.CSTime,
		&stats.Priority,
		&stats.Nice,
		&stats.NumThreads,
		&ignore,
		&stats.Starttime,
		&stats.VSize,
		&stats.RSS,
	)

	if err != nil {
		return ProcStat{}
	}

	return stats
}

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
	filepath := fmt.Sprint("/proc/", strconv.Itoa(pid), "/schedstat")
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

// GetCmdLine reads the cmdline for a process from /proc
func GetCmdLine(pid int) string {
	filepath := fmt.Sprint("/proc/", strconv.Itoa(pid), "/cmdline")
	filecontent, _ := ioutil.ReadFile(filepath)
	return string(filecontent)
}
