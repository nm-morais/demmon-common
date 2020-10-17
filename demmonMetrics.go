package demmonMetrics

import (
	"io"

	"github.com/nm-morais/demmon-common/metrics"
)

var defaultSet = metrics.NewSet()

func WriteMetrics(w io.Writer) {
	defaultSet.WriteMetrics(w)
}

// NewCounter registers and returns new counter with the given name.
// The returned counter is safe to use from concurrent goroutines.
func NewCounter(name string) *metrics.Counter {
	return defaultSet.NewCounter(name)
}

// NewGauge registers and returns gauge with the given name, which calls f
// to obtain gauge value.
func NewGauge(name string, f func() float64) *metrics.Gauge {
	return defaultSet.NewGauge(name, f)
}

func NewHistogram(name string) *metrics.Histogram {
	return defaultSet.NewHistogram(name)
}
