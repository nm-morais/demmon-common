package metrics

import (
	"fmt"
	"io"
)

// Gauge is a float64 gauge.
//
// See also Counter, which could be used as a gauge with Set and Dec calls.
type Gauge struct {
	name string
	f    func() float64
}

// Get returns the current value for g.
func (g *Gauge) Get() float64 {
	return g.f()
}

func (g *Gauge) Name() string {
	return g.name
}

func (g *Gauge) MarshalTo(prefix string, w io.Writer) {
	v := g.f()
	if float64(int64(v)) == v {
		// Marshal integer values without scientific notation
		fmt.Fprintf(w, "%s %d\n", prefix, int64(v))
	} else {
		fmt.Fprintf(w, "%s %g\n", prefix, v)
	}
}
