package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"kvmtop/config"
)

// ProcPIDNetDev represents entries of /proc/<pid>/net/dev file
type ProcPIDNetDev struct {
	// The process ID.
	PID int

	Dev string

	ReceivedBytes      uint64
	ReceivedPackets    uint64
	ReceivedErrs       uint64
	ReceivedDrop       uint64
	ReceivedFifo       uint64
	ReceivedFrame      uint64
	ReceivedCompressed uint64
	ReceivedMulticast  uint64

	TransmittedBytes      uint64
	TransmittedPackets    uint64
	TransmittedErrs       uint64
	TransmittedDrop       uint64
	TransmittedFifo       uint64
	TransmittedColls      uint64
	TransmittedCarrier    uint64
	TransmittedCompressed uint64
}

// GetProcPIDNetDev reads the net/dev file for given pid and device name from procfs
func GetProcPIDNetDev(pid int, dev string) ProcPIDNetDev {
	stats := ProcPIDNetDev{PID: pid, Dev: dev}

	filepath := fmt.Sprint(config.Options.ProcFS, "/net/dev")
	if pid != 0 {
		filepath = fmt.Sprint(config.Options.ProcFS, "/", strconv.Itoa(pid), "/net/dev")
	}

	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	format := "" + dev + ": %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d"

	foundDevStats := false
	for scanner.Scan() {
		row := strings.Trim(scanner.Text(), " ")
		if !strings.Contains(row, dev+":") {
			// only parse line with specified device
			continue
		}
		foundDevStats = true
		_, err := fmt.Sscanf(
			string(row), format,
			&stats.ReceivedBytes,
			&stats.ReceivedPackets,
			&stats.ReceivedErrs,
			&stats.ReceivedDrop,
			&stats.ReceivedFifo,
			&stats.ReceivedFrame,
			&stats.ReceivedCompressed,
			&stats.ReceivedMulticast,

			&stats.TransmittedBytes,
			&stats.TransmittedPackets,
			&stats.TransmittedErrs,
			&stats.TransmittedDrop,
			&stats.TransmittedFifo,
			&stats.TransmittedColls,
			&stats.TransmittedCarrier,
			&stats.TransmittedCompressed,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse row in proc net/dev: %s\n", err)
			return ProcPIDNetDev{}
		}
	}

	if !foundDevStats {
		fmt.Fprintf(os.Stderr, "could not find network device %s\n", dev)
		return ProcPIDNetDev{}
	}

	return stats
}
