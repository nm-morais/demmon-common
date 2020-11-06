package exporters

import (
	"github.com/nm-morais/demmon-common/timeseries"
)

type Exporter interface {
	Name() string
	Get() timeseries.Value
}
