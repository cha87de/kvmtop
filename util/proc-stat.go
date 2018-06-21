package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/cha87de/kvmtop/config"
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
	VSize uint64
	// Resident set size in pages.
	RSS int

	// Current soft limit in bytes on the rss of the process; see the description of RLIMIT_RSS in	getrlimit(2).
	RSSLim uint64
	// The address above which program text can run.
	Startcode uint64
	// The address below which program text can run.
	Endcode uint64
	// The address of the start (i.e., bottom) of the stack.
	Startstack uint64
	// The current value of ESP (stack pointer), as found in the kernel stack page for the process.
	Kstkesp uint64
	// The current EIP (instruction pointer).
	Kstkeip uint64
	// The bitmap of pending signals, displayed as a deci‐
	// mal number.  Obsolete, because it does not provide
	// information on real-time signals; use
	// /proc/[pid]/status instead.
	Signal uint64
	// The bitmap of blocked signals, displayed as a deci‐
	// mal number.  Obsolete, because it does not provide
	// information on real-time signals; use
	// /proc/[pid]/status instead.
	Blocked uint64
	// The bitmap of ignored signals, displayed as a deci‐
	// mal number.  Obsolete, because it does not provide
	// information on real-time signals; use
	// /proc/[pid]/status instead.
	Sigignore uint64
	// The bitmap of caught signals, displayed as a decimal
	// number.  Obsolete, because it does not provide
	// information on real-time signals; use
	// /proc/[pid]/status instead.
	Sigcatch uint64
	// This is the "channel" in which the process is wait‐
	// ing.  It is the address of a location in the kernel
	// where the process is sleeping.  The corresponding
	// symbolic name can be found in /proc/[pid]/wchan.
	Wchan uint64
	// Number of pages swapped (not maintained).
	Nswap uint64
	// Cumulative nswap for child processes (not main‐
	// tained).
	Cnswap uint64
	// Signal to be sent to parent when we die.
	ExitSignal int
	// CPU number last executed on.
	Processor int
	// Real-time scheduling priority, a number in the range
	// 1 to 99 for processes scheduled under a real-time
	// policy, or 0, for non-real-time processes (see
	// sched_setscheduler(2)).
	RtPriority uint
	// Scheduling policy (see sched_setscheduler(2)).
	// Decode using the SCHED_* constants in linux/sched.h.
	Policy uint
	// Aggregated block I/O delays, measured in clock ticks
	// (centiseconds).
	DelayacctBlkioTicks uint64
	// Guest time of the process (time spent running a vir‐
	// tual CPU for a guest operating system), measured in
	// clock ticks (divide by sysconf(_SC_CLK_TCK)).
	GuestTime uint64
	// Guest time of the process's children, measured in
	// clock ticks (divide by sysconf(_SC_CLK_TCK)).
	CGuestTime uint64
	// Address above which program initialized and unini‐
	// tialized (BSS) data are placed.
	StartData uint64
	// Address below which program initialized and unini‐
	// tialized (BSS) data are placed.
	EndData uint64
	// Address above which program heap can be expanded
	// with brk(2).
	StartBrk uint64
	// Address above which program command-line arguments
	// (argv) are placed.
	ArgStart uint64
	// Address below program command-line arguments (argv)
	// are placed.
	ArgEnd uint64
	// Address above which program environment is placed.
	EnvStart uint64
	// Address below which program environment is placed.
	EnvEnd uint64
	// The thread's exit status in the form reported by
	// waitpid(2).
	ExitCode uint64
}

// GetProcStat reads and returns the stat for a process from the proc fs
func GetProcStat(pid int) ProcStat {
	stats := ProcStat{PID: pid}
	filepath := fmt.Sprint(config.Options.ProcFS, "/", strconv.Itoa(pid), "/stat")
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
		&stats.RSSLim,
		&stats.Startcode,
		&stats.Endcode,
		&stats.Startstack,
		&stats.Kstkesp,
		&stats.Kstkeip,
		&stats.Signal,
		&stats.Blocked,
		&stats.Sigignore,
		&stats.Sigcatch,
		&stats.Wchan,
		&stats.Nswap,
		&stats.Cnswap,
		&stats.ExitSignal,
		&stats.Processor,
		&stats.RtPriority,
		&stats.Policy,
		&stats.DelayacctBlkioTicks,
		&stats.GuestTime,
		&stats.CGuestTime,
		&stats.StartData,
		&stats.EndData,
		&stats.StartBrk,
		&stats.ArgStart,
		&stats.ArgEnd,
		&stats.EnvStart,
		&stats.EnvEnd,
		&stats.ExitCode,
	)

	if err != nil {
		return ProcStat{}
	}

	return stats
}
