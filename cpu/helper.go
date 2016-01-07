package cpu

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
	"github.com/cha87de/kvmtop/models"
)

type VmCpuUtilisation struct{
	Timestamp time.Time
	Avg VmCoreUtilisation
	Avg_other float64
	VCores []VmCoreUtilisation
}

type VmCoreUtilisation struct{
	Inside float64
	Outside float64
	Steal float64
}

type VmCpuSchedStat struct{
	vCores []CpuSchedStat
	other CpuSchedStat
}

type CpuSchedStat struct{
	timestamp time.Time
	cputime time.Duration
	runqueue time.Duration
	//timeslices time.Duration
}

func (first *VmCpuSchedStat) calculateDiff(second VmCpuSchedStat) VmCpuUtilisation {
	vmCpuUtilisation := VmCpuUtilisation{}
	vmCoreUtilisationAverage := VmCoreUtilisation{0,0,0}
	
	// run through each core
	for i, vCoreFirst := range first.vCores{
		vCoreSecond := second.vCores[i]
		vmCoreUtilisation := VmCoreUtilisation{0,0,0}
		
		var (
			firstMeasurement float64
			secondMeasurement float64
		)

		timeDiff := (vCoreSecond.timestamp.Sub(vCoreFirst.timestamp).Seconds())

		// outside = cputime
		firstMeasurement = vCoreFirst.cputime.Seconds()
		secondMeasurement = vCoreSecond.cputime.Seconds()
		vmCoreUtilisation.Outside = (secondMeasurement - firstMeasurement) / timeDiff 
		vmCoreUtilisationAverage.Outside = vmCoreUtilisationAverage.Outside + vmCoreUtilisation.Outside
				
		// steal = runqueue
		firstMeasurement = vCoreFirst.runqueue.Seconds()
		secondMeasurement = vCoreSecond.runqueue.Seconds()
		vmCoreUtilisation.Steal = (secondMeasurement - firstMeasurement) / timeDiff 
		vmCoreUtilisationAverage.Steal = vmCoreUtilisationAverage.Steal + vmCoreUtilisation.Steal

		// inside = cputime+runqueue
		//firstMeasurement = vCoreFirst.cputime.Seconds() + vCoreFirst.runqueue.Seconds()
		//secondMeasurement = vCoreSecond.cputime.Seconds() + vCoreSecond.runqueue.Seconds()
		//vmCoreUtilisation.Inside = (secondMeasurement - firstMeasurement) / timeDiff
		vmCoreUtilisation.Inside = vmCoreUtilisation.Outside + vmCoreUtilisation.Steal
		vmCoreUtilisationAverage.Inside = vmCoreUtilisationAverage.Inside + vmCoreUtilisation.Inside
		
		vmCpuUtilisation.VCores = append(vmCpuUtilisation.VCores, vmCoreUtilisation)
		vmCpuUtilisation.Timestamp = vCoreSecond.timestamp
	}
	
	// calculate average over all cores
	vmCoreUtilisationAverage.Inside = vmCoreUtilisationAverage.Inside / float64(len(first.vCores))
	vmCoreUtilisationAverage.Outside = vmCoreUtilisationAverage.Outside / float64(len(first.vCores))
	vmCoreUtilisationAverage.Steal = vmCoreUtilisationAverage.Steal / float64(len(first.vCores))
	vmCpuUtilisation.Avg = vmCoreUtilisationAverage

	// calculate "other" utilisation field
	timeDiff := (second.other.timestamp.Sub(first.other.timestamp).Seconds())
	secondMeasurement := second.other.cputime.Seconds() + second.other.runqueue.Seconds()
	firstMeasurement := first.other.cputime.Seconds() + first.other.runqueue.Seconds()
	vmCpuUtilisation.Avg_other = (secondMeasurement - firstMeasurement) / timeDiff
	
	return vmCpuUtilisation
}

func lookupStats(vm models.VirtualMachine) (VmCpuSchedStat, error) {
	cpuSchedStats := []CpuSchedStat{}
	for _, taskId := range vm.VCpuTasks() {
		cpuSchedStat, err := getCpuSchedStat(taskId)
		if err != nil{
			return VmCpuSchedStat{}, err
		}
		cpuSchedStats = append(cpuSchedStats, cpuSchedStat)
	}
	other := CpuSchedStat{}
	// TODO fill other
	//getTasks(vm)
	return VmCpuSchedStat{cpuSchedStats, other}, nil 
}

func getCpuSchedStat(taskid string) (CpuSchedStat, error) {
        schedstat_file := fmt.Sprint("/proc/", taskid, "/schedstat")
        schedstat_content, err := ioutil.ReadFile(schedstat_file)
		if err != nil {
			return CpuSchedStat{}, err
		}
	    schedstat := strings.Split(string(schedstat_content), "\n")
	    schedstat_val := strings.Split(schedstat[0], " ")
	    cputime, _ := strconv.ParseInt(schedstat_val[0], 10, 64)
	    runqueue, _ := strconv.ParseInt(schedstat_val[1], 10, 64)
	    
	    return CpuSchedStat{time.Now(), 
	    	time.Duration(cputime), 
	    	time.Duration(runqueue)}, nil
}

func getTasks(vm models.VirtualMachine) {
		folder := fmt.Sprint("/proc/", vm.Ppid(), "/task/")
		files, _ := ioutil.ReadDir(folder)
	    for _, f := range files {
	           taskId := f.Name()
	           fmt.Printf("%s ", taskId)
	    }
}
