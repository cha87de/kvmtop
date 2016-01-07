package models

import (
	"time"
	"strconv"
)

type MeasurementItem struct {
		Name string
}

type Statistic struct {
		Timestamp time.Time
		Values map[string] string
		Eval func(Statistic) int64
}

func (stats *Statistic) GetValueAsInt(key string) int64 {
	val64, err := strconv.ParseInt(stats.Values[key], 10, 64)
	if err != nil {
		return 0
	}
	return val64
}

func (statsFirst Statistic) DiffPerTime(statsSecond Statistic) int64 {
	diffValue := statsSecond.Eval(statsSecond)-statsFirst.Eval(statsFirst)
	diffTime := int64(statsSecond.Timestamp.Sub(statsFirst.Timestamp).Seconds())
	if diffTime == 0 {
		return 0
	}
	diffPerTime := int64(diffValue / diffTime)
	return diffPerTime
}

func (statsFirst *Statistic) DiffPerTimeField(statsSecond Statistic, fieldIndex string) int64 {
	diffValue := statsSecond.GetValueAsInt(fieldIndex)-statsFirst.GetValueAsInt(fieldIndex)
	diffTime := int64(statsSecond.Timestamp.Sub(statsFirst.Timestamp).Seconds())
	if diffTime == 0 {
		return 0
	}
	diffPerTime := int64(diffValue / diffTime)
	return diffPerTime
}