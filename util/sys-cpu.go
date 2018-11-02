package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SysCPU reflects cpu system information from /sys/devices/system/cpu/cpu*
type SysCPU struct {
	MaxFreq float32
	MinFreq float32
	CurFreq float32
}

// GetSysCPU returns the system CPU information for available cores
func GetSysCPU() []SysCPU {
	stats := []SysCPU{}

	files, err := filepath.Glob("/sys/devices/system/cpu/cpu[0-999]")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read cpu infos from sys fs: %s\n", err)
		return stats
	}

	var filepath string
	var filecontent []byte
	for _, f := range files {
		cpuStat := SysCPU{}

		filepath = fmt.Sprint(f + "/cpufreq/cpuinfo_max_freq")
		filecontent, _ = ioutil.ReadFile(filepath)
		fmt.Fscan(
			bytes.NewBuffer(filecontent),
			&cpuStat.MaxFreq,
		)

		filepath = fmt.Sprint(f + "/cpufreq/cpuinfo_min_freq")
		filecontent, _ = ioutil.ReadFile(filepath)
		fmt.Fscan(
			bytes.NewBuffer(filecontent),
			&cpuStat.MinFreq,
		)

		filepath = fmt.Sprint(f + "/cpufreq/scaling_cur_freq")
		filecontent, _ = ioutil.ReadFile(filepath)
		fmt.Fscan(
			bytes.NewBuffer(filecontent),
			&cpuStat.CurFreq,
		)

		stats = append(stats, cpuStat)

	}

	return stats
}
