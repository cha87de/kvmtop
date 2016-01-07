package util

import (
	"github.com/cha87de/kvmtop/models"
	"os/exec"
	"bytes"
	"strings"
	"time"
	"log"
)

func VirshXList(command string, vmname string) ([]models.MeasurementItem, error) {
	cmd := exec.Command("virsh", "--connect=qemu:///system", command, vmname)
	var domlist bytes.Buffer
	cmd.Stdout = &domlist
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	domlistArr := strings.Split(domlist.String(), "\n")
	domlistArr = domlistArr[2:] // cut 2 lines for header
	var itemList []models.MeasurementItem
	for _, line := range domlistArr {
		if(line == ""){
			continue
		}
		lineArr := strings.Split(line, " ")
		item := models.MeasurementItem{lineArr[0]}
		itemList = append(itemList, item)
	}
	return itemList, nil	
	
}

func VirshXDetails(command string, vmname string, itemName string, indexKey int, indexValue int, evalFunc func(models.Statistic) int64) (models.Statistic, error){
	
	var cmd *exec.Cmd
	if itemName != "" {
		cmd = exec.Command("virsh", "--connect=qemu:///system", command, vmname, itemName)
	}else{
		cmd = exec.Command("virsh", "--connect=qemu:///system", command, vmname)
	}
	var domifstat bytes.Buffer
	cmd.Stdout = &domifstat
	err := cmd.Run()
	if err != nil {
		log.Printf("Error while exec virsh %a", err)
		return models.Statistic{}, err
	}
	domifstatArr := strings.Split(domifstat.String(), "\n")
	ifstats := make(map[string]string)
	for _, line := range domifstatArr {
		if(line == ""){
			continue
		}
		lineArr := strings.Fields(line)
//		val64, err := strconv.ParseInt(lineArr[indexValue], 10, 64)
//		if err != nil {
//			continue;
//		}
		val := lineArr[indexValue]
		index := strings.Trim(lineArr[indexKey], ":")
		ifstats[index] = val
	}
	stat := models.Statistic{time.Now(), ifstats, evalFunc}
	return stat, nil
	
}

