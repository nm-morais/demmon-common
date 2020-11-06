package default_exporters

import (
	"sync"

	"github.com/nm-morais/demmon-common/default_plugin"
	"github.com/nm-morais/demmon-common/timeseries"
)

// FloatGauge is a float64 gauge.
//
// See also Counter, which could be used as a gauge with Set and Dec calls.
type FloatGauge struct {
	name string
	mu   *sync.Mutex
	v    *default_plugin.FloatValue
}

func NewFloatGauge(name string, initialVal float64) *FloatGauge {
	return &FloatGauge{
		name: name,
		mu:   &sync.Mutex{},
		v: &default_plugin.FloatValue{
			V: initialVal,
		},
	}
}

// Get returns the current value for g.
func (g *FloatGauge) Get() timeseries.Value {
	g.mu.Lock()
	v := g.v
	g.mu.Unlock()
	return v
}

// Set sets fc value to n.
func (fc *FloatGauge) Set(n float64) {
	fc.mu.Lock()
	fc.v.V = n
	fc.mu.Unlock()
}

func (g *FloatGauge) Name() string {
	return g.name
}

// Counter is a timeseries.Value counter guarded by RWmutex.
//
// It may be used as a gauge if Add and Sub are called.
type FloatCounter struct {
	name string
	mu   *sync.Mutex
	v    *default_plugin.FloatValue
}

func NewCounter(name string, initialValue float64) *FloatCounter {
	return &FloatCounter{
		name: name,
		mu:   &sync.Mutex{},
		v: &default_plugin.FloatValue{
			V: initialValue,
		},
	}
}

// Add adds n to fc.
func (fc *FloatCounter) Add(n float64) {
	fc.mu.Lock()
	fc.v.V += n
	fc.mu.Unlock()
}

// Sub substracts n from fc.
func (fc *FloatCounter) Sub(n float64) {
	fc.mu.Lock()
	fc.v.V -= n
	fc.mu.Unlock()
}

// Set sets fc value to n.
func (fc *FloatCounter) Set(n float64) {
	fc.mu.Lock()
	fc.v.V = n
	fc.mu.Unlock()
}

// Get returns the current value for fc.
func (fc *FloatCounter) Get() timeseries.Value {
	fc.mu.Lock()
	v := fc.v
	fc.mu.Unlock()
	return v
}

func (fc *FloatCounter) Name() string {
	return fc.name
}
