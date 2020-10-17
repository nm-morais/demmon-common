package metrics

import (
	"fmt"
	"io"
	"sync"
)

// Counter is a float64 counter guarded by RWmutex.
//
// It may be used as a gauge if Add and Sub are called.
type Counter struct {
	name string
	mu   sync.Mutex
	n    float64
}

// Add adds n to fc.
func (fc *Counter) Add(n float64) {
	fc.mu.Lock()
	fc.n += n
	fc.mu.Unlock()
}

// Sub substracts n from fc.
func (fc *Counter) Sub(n float64) {
	fc.mu.Lock()
	fc.n -= n
	fc.mu.Unlock()
}

// Get returns the current value for fc.
func (fc *Counter) Get() float64 {
	fc.mu.Lock()
	n := fc.n
	fc.mu.Unlock()
	return n
}

// Set sets fc value to n.
func (fc *Counter) Set(n float64) {
	fc.mu.Lock()
	fc.n = n
	fc.mu.Unlock()
}

func (fc *Counter) Name() string {
	return fc.name
}

// marshalTo marshals fc with the given prefix to w.
func (fc *Counter) MarshalTo(prefix string, w io.Writer) {
	v := fc.Get()
	fmt.Fprintf(w, "%s %g\n", prefix, v)
}
