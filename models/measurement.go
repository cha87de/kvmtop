package models

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Measurement represents one measurement value at a timestamp
type Measurement struct {
	Value     []byte
	Timestamp time.Time
}

// CreateMeasurement creates a new measurement with recent time and bytes from provided value
func CreateMeasurement(value interface{}) Measurement {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(value)
	if err != nil {
		return Measurement{}
	}
	return Measurement{
		Value:     buffer.Bytes(),
		Timestamp: time.Now(),
	}
}
