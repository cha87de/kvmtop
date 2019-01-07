package config

import "time"

// ProfilerOptionsType describes profiler options
type ProfilerOptionsType struct {
	States        int           `long:"states" default:"4"`
	BufferSize    int           `long:"buffersize" default:"10"`
	History       int           `long:"history" default:"1"`
	FilterStdDevs int           `long:"filterstddevs" default:"-1"`
	OutputFreq    time.Duration `long:"outputFreq" default:"60"`
}
