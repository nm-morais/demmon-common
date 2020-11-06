package default_plugin

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/nm-morais/demmon-common/timeseries"
)

const PluginName = "Default_plugin"

func NewFloatVal() *FloatValue {
	return &FloatValue{
		V: 0,
	}
}

type FloatValue struct {
	V float64
}

func (f *FloatValue) Val() interface{} {
	return f.V
}

func MarshalFloatValue(vGeneric timeseries.Value) ([]byte, error) {
	bytes := make([]byte, 8)
	v, ok := vGeneric.(*FloatValue)
	if !ok {
		return nil, fmt.Errorf("incorrect type %+v, should be floatValue", vGeneric)
	}
	binary.BigEndian.PutUint64(bytes, math.Float64bits(v.V))
	return bytes, nil
}

func UnmarshalFloatValue(b []byte) (timeseries.Value, error) {
	return &FloatValue{math.Float64frombits(binary.BigEndian.Uint64(b))}, nil
}
