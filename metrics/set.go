package metrics

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"sync"
)

type namedMetric struct {
	name   string
	metric metric
}

type metric interface {
	Name() string
	MarshalTo(prefix string, w io.Writer)
}

// Set is a set of metrics.
//
// Metrics belonging to a set are exported separately from global metrics.
//
// Set.WriteMetrics must be called for exporting metrics from the set.
type Set struct {
	mu sync.Mutex
	a  []*namedMetric
	m  map[string]*namedMetric
}

// NewSet creates new set of metrics.
func NewSet() *Set {
	return &Set{
		m: make(map[string]*namedMetric),
	}
}

// WriteMetrics writes all the metrics from s to w
func (s *Set) WriteMetrics(w io.Writer) {
	// Collect all the metrics in in-memory buffer in order to prevent from long locking due to slow w.
	var bb bytes.Buffer
	lessFunc := func(i, j int) bool {
		return s.a[i].name < s.a[j].name
	}
	s.mu.Lock()
	if !sort.SliceIsSorted(s.a, lessFunc) {
		sort.Slice(s.a, lessFunc)
	}
	sa := append([]*namedMetric(nil), s.a...)
	s.mu.Unlock()

	// Call marshalTo without the global lock, since certain metric types such as Gauge
	// can call a callback, which, in turn, can try calling s.mu.Lock again.
	for _, nm := range sa {
		nm.metric.MarshalTo(nm.name, &bb)
	}
	w.Write(bb.Bytes())
}

func (s *Set) NewHistogram(name string) *Histogram {
	h := &Histogram{
		name: name,
	}
	s.registerMetric(name, h)
	return h
}

func (s *Set) NewCounter(name string) *Counter {
	c := &Counter{
		name: name,
	}
	s.registerMetric(name, c)
	return c
}

func (s *Set) NewGauge(name string, f func() float64) *Gauge {
	if f == nil {
		panic(fmt.Errorf("BUG: f cannot be nil"))
	}
	g := &Gauge{
		name: name,
		f:    f,
	}
	s.registerMetric(name, g)
	return g
}

func (s *Set) registerMetric(name string, m metric) {
	s.mu.Lock()
	// defer will unlock in case of panic
	// checks in test
	defer s.mu.Unlock()
	s.mustRegisterLocked(name, m)
}

// mustRegisterLocked registers given metric with
// the given name. Panics if the given name was
// already registered before.
func (s *Set) mustRegisterLocked(name string, m metric) {
	nm, ok := s.m[name]
	if !ok {
		nm = &namedMetric{
			name:   name,
			metric: m,
		}
		s.m[name] = nm
		s.a = append(s.a, nm)
	}
	if ok {
		panic(fmt.Errorf("BUG: metric %q is already registered", name))
	}
}

// UnregisterMetric removes metric with the given name from s.
//
// True is returned if the metric has been removed.
// False is returned if the given metric is missing in s.
func (s *Set) UnregisterMetric(name string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.m[name]
	if !ok {
		return false
	}
	delete(s.m, name)

	deleteFromList := func(metricName string) {
		for i, nm := range s.a {
			if nm.name == metricName {
				s.a = append(s.a[:i], s.a[i+1:]...)
				return
			}
		}
		panic(fmt.Errorf("BUG: cannot find metric %q in the list of registered metrics", name))
	}
	// remove metric from s.a
	deleteFromList(name)
	return true
}

// ListMetricNames returns a list of all the metrics in s.
func (s *Set) ListMetricNames() []string {
	var list []string
	for name := range s.m {
		list = append(list, name)
	}
	return list
}
