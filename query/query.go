package query

import (
	"github.com/nm-morais/demmon-common/timeseries"
)

type MetricMetadata struct {
}

type GlobalAggregationFunc struct {
	QueryInput          string // this value is used to perform a prefix search and fetch all the timeseries which are used as input to the aggregation function
	AggregationFunc     func(...timeseries.TimeSeries) []byte
	ResultingMetricName string // the name of the resulting metric
}

type Alert struct {
	Query string
}
