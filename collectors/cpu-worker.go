package collectors

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"regexp"
	"strconv"

	"fmt"

	"github.com/cha87de/kvmtop/models"
	libvirt "github.com/libvirt/libvirt-go"
)

var regSchedStat *regexp.Regexp

func cpuLookup(domain *models.Domain, libvirtDomain libvirt.Domain) {
	// get amount of cores
	vcpus, err := libvirtDomain.GetVcpus()
	if err != nil {
		return
	}
	cores := len(vcpus)
	newMeasurementCores := models.CreateMeasurement(cores)
	domain.AddMetricMeasurement("cpu_cores", newMeasurementCores)

	// get core thread IDs
	vCPUThreads, err := libvirtDomain.QemuMonitorCommand("info cpus", libvirt.DOMAIN_QEMU_MONITOR_COMMAND_HMP)
	if err != nil {
		return
	}
	regThreadID := regexp.MustCompile("thread_id=([0-9]*)\\s")
	threadIDsRaw := regThreadID.FindAllStringSubmatch(vCPUThreads, -1)
	threadIDs := make([]int, len(threadIDsRaw))
	for i, thread := range threadIDsRaw {
		threadIDs[i], _ = strconv.Atoi(thread[1])
	}
	newMeasurementThreads := models.CreateMeasurement(threadIDs)
	domain.AddMetricMeasurement("cpu_threadIDs", newMeasurementThreads)
}

func cpuCollect(domain *models.Domain) {
	// get threadIDs
	var threadIDs []int
	if metric, ok := domain.Metrics["cpu_threadIDs"]; ok {
		if len(metric.Values) > 0 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			dec := gob.NewDecoder(reader)
			dec.Decode(&threadIDs)
		}
	}

	// get statistics for threadIDs
	if regSchedStat == nil {
		regSchedStat = regexp.MustCompile("[0-9]*")
	}
	var cputimes []int64
	var runqueues []int64
	for _, threadID := range threadIDs {
		schedstatFile := fmt.Sprint("/proc/", threadID, "/schedstat")
		schedstatFileContent, _ := ioutil.ReadFile(schedstatFile)
		schedStatCounters := regSchedStat.FindAllStringSubmatch(string(schedstatFileContent), -1)
		if len(schedStatCounters) < 2 {
			// schedstat file unreadable
			continue
		}
		cputime, _ := strconv.ParseInt(schedStatCounters[0][0], 10, 64)
		runqueue, _ := strconv.ParseInt(schedStatCounters[1][0], 10, 64)
		cputimes = append(cputimes, cputime)
		runqueues = append(runqueues, runqueue)
	}

	// store cputimes and runqueues as metrics
	newMeasurementCputimes := models.CreateMeasurement(cputimes)
	newMeasurementRunqueues := models.CreateMeasurement(runqueues)
	domain.AddMetricMeasurement("cpu_times", newMeasurementCputimes)
	domain.AddMetricMeasurement("cpu_runqueues", newMeasurementRunqueues)
}

func cpuPrint(domain *models.Domain) []string {
	var cores string
	if metric, ok := domain.Metrics["cpu_cores"]; ok {
		if len(metric.Values) > 0 {
			byteValue := metric.Values[0].Value
			reader := bytes.NewReader(byteValue)
			dec := gob.NewDecoder(reader)

			var coresRaw int
			dec.Decode(&coresRaw)
			cores = fmt.Sprintf("%d", coresRaw)
		}
	}

	var cputimes []string
	var cputime string
	if metric, ok := domain.Metrics["cpu_times"]; ok {
		if len(metric.Values) > 1 {
			byteValue1 := metric.Values[0].Value
			reader1 := bytes.NewReader(byteValue1)
			dec1 := gob.NewDecoder(reader1)

			byteValue2 := metric.Values[1].Value
			reader2 := bytes.NewReader(byteValue2)
			dec2 := gob.NewDecoder(reader2)

			var cputimesRaw1 []int64
			var cputimesRaw2 []int64
			dec1.Decode(&cputimesRaw1)
			dec2.Decode(&cputimesRaw2)

			timeDiff := metric.Values[0].Timestamp.Sub(metric.Values[1].Timestamp).Seconds()
			timeConversionFactor := 1000000000 / timeDiff

			// for each core ...
			var cputimeSum float64
			for i, cputime1 := range cputimesRaw1 {
				if len(cputimesRaw2) <= i {
					continue
				}
				cputime2 := cputimesRaw2[i]
				cputime := float64(cputime1-cputime2) / timeConversionFactor
				cputimeSum = cputimeSum + cputime
				cputimes = append(cputimes, fmt.Sprintf("%.0f", cputime*100))
			}

			cputime = fmt.Sprintf("%.0f", cputimeSum/float64(len(cputimesRaw1))*100)
		}
	}

	result := append([]string{cores}, cputime)
	result = append(result, cputimes[0:]...)
	return result
}
