package demmonMetrics

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestHistogram(t *testing.T) {
	// Define a histogram in global scope.
	var h = NewHistogram(`request_duration_seconds_path_/foo/bar`)

	startTime := time.Now()

	h.UpdateDuration(startTime)

	buf := &bytes.Buffer{}
	h.MarshalTo("cenas", buf)

	fmt.Println(string(buf.Bytes()))

	t.FailNow()
}

func TestCounter(t *testing.T) {
	// Define a histogram in global scope.
	var c = NewCounter(`myCounter`)
	c.Add(1)
	c.Add(1)

	buf := &bytes.Buffer{}
	c.MarshalTo("cenas", buf)

	fmt.Println(string(buf.Bytes()))

	t.FailNow()
}

func TestGauge(t *testing.T) {
	// Define a histogram in global scope.
	var c = NewGauge(`myGauge`, func() float64 { return 1 })

	c.Get()

	buf := &bytes.Buffer{}
	c.MarshalTo("myGauge", buf)

	fmt.Println(string(buf.Bytes()))
}
