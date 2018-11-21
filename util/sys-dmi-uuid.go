package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// SysDmiUUID reflects the systems UUID from /sys/devices/virtual/dmi/id/product_uuid
type SysDmiUUID struct {
	Value string
}

// GetSysDmiUUID reads the hosts DMI UUID from /sys/devices/virtual/dmi/id/product_uuid
// Please Note: root privileges are required to read the file
func GetSysDmiUUID() SysDmiUUID {
	stat := SysDmiUUID{}
	filepath := fmt.Sprint("/sys/devices/virtual/dmi/id/product_uuid")
	filecontent, _ := ioutil.ReadFile(filepath)
	fmt.Fscan(
		bytes.NewBuffer(filecontent),
		&stat.Value,
	)
	return stat
}
