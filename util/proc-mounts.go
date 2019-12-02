package util

import (
	"bufio"
	"fmt"
	"os"

	"kvmtop/config"
)

// ProcMount describes one entry (row) of /proc/mounts
type ProcMount struct {
	Device         string
	Mountpoint     string
	FileSystemType string
	Options        string
	dummy1         int
	dummy2         int
}

// GetProcMounts reads and returns an array of mount points defined in /proc/mounts
func GetProcMounts() []ProcMount {
	mounts := []ProcMount{}
	filepath := fmt.Sprint(config.Options.ProcFS, "/mounts")

	file, _ := os.Open(filepath)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	format := "%s %s %s %s %d %d"

	for scanner.Scan() {
		row := scanner.Text()
		mount := ProcMount{}

		_, err := fmt.Sscanf(
			string(row), format,
			&mount.Device,
			&mount.Mountpoint,
			&mount.FileSystemType,
			&mount.Options,
			&mount.dummy1,
			&mount.dummy2,
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot parse row in proc mounts: %s\n", err)
			continue
		}
		mounts = append(mounts, mount)
	}

	return mounts
}
