package timeseries

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

type Granularity struct {
	Granularity time.Duration
	Count       int
}

var DefaultGranularities = []Granularity{
	{time.Second, 60},
	{time.Minute, 60},
	{time.Hour, 24},
}

type ExportConf struct {
	EmitFrequency time.Duration
	BatchMetric   bool
	Granularity   Granularity
}

type Value interface{}

type PointValue struct {
	TS    time.Time
	Value Value
}

type TimeSeries interface {
	All() []*PointValue
	AddPoint(p *PointValue)
	Last() (*PointValue, error)
	// MergeWith(series Timeseries) error
	Range(start time.Time, end time.Time) ([]*PointValue, error)
	sync.Locker
}

type UnmarshalFunc = func([]byte) (Value, error)
type MarshalFunc = func(Value) ([]byte, error)
type AggregationFunc = func(...TimeSeries) Value

func MarshalPVAsTextLine(service, name, origin string, pv *PointValue, encodeFunc MarshalFunc) (string, error) {
	val, err := encodeFunc(pv.Value)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s/%s %s %d", service, name, origin, base64.StdEncoding.EncodeToString(val), pv.TS.UnixNano()), nil
}

func MarshalPVArrAsText(service, name, origin string, arr []*PointValue, encodeFunc MarshalFunc) (string, error) {
	out := ""
	for _, v := range arr {
		v, err := MarshalPVAsTextLine(service, name, origin, v, encodeFunc)
		if err != nil {
			return "", err
		}
		out += string(v)
	}
	return out, nil
}

func MarshalPVArrAsStrArr(service, name, origin string, arr []*PointValue, encodeFunc MarshalFunc) ([]string, error) {
	res := make([]string, len(arr))
	for idx, v := range arr {
		v, err := MarshalPVAsTextLine(service, name, origin, v, encodeFunc)
		if err != nil {
			return nil, err
		}
		res[idx] = v
	}
	return res, nil
}
