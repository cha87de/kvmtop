package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/cha87de/kvmtop/config"
)

// ProcPIDIO defines the fields of a /proc/[pid]/io file
// cf. https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/filesystems/proc.txt?id=HEAD
type ProcPIDIO struct {
	// The process ID.
	PID int
	//  number of bytes the process read, using any read-like system call (from files, pipes, tty...).
	Rchar uint64
	// number of bytes the process wrote using any write-like system call.
	Wchar uint64
	// number of read-like system call invocations that the process performed.
	Syscr uint64
	// number of write-like system call invocations that the process performed.
	Syscw uint64
	// number of bytes the process directly read from disk.
	Read_bytes uint64
	// number of bytes the process originally dirtied in the page-cache (assuming they will go to disk later).
	Write_bytes uint64
	// number of bytes the process "un-dirtied" - e.g. using an "ftruncate" call that truncated pages from the page-cache.
	Cancelled_write_bytes uint64
}

// GetProcPIDIO reads and returns the io for a process from the proc fs
func GetProcPIDIO(pid int) ProcPIDIO {
	stats := ProcPIDIO{PID: pid}
	filepath := fmt.Sprint(config.Options.ProcFS, "/", strconv.Itoa(pid), "/io")
	filecontent, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read proc io: %s\n", err)
		return ProcPIDIO{}
	}

	ioFormat := "rchar: %d\nwchar: %d\nsyscr: %d\nsyscw: %d\n" +
		"read_bytes: %d\nwrite_bytes: %d\n" +
		"cancelled_write_bytes: %d\n"

	_, err = fmt.Sscanf(
		string(filecontent), ioFormat,
		&stats.Rchar,
		&stats.Wchar,
		&stats.Syscr,
		&stats.Syscw,
		&stats.Read_bytes,
		&stats.Write_bytes,
		&stats.Cancelled_write_bytes,
	)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse proc io: %s\n", err)
		return ProcPIDIO{}
	}

	return stats
}
