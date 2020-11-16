package routes

import "fmt"

const (
	Dial = "/dial"
)

type RequestType int

func NewRequest(reqNr int) RequestType {
	return RequestType(reqNr)
}

const (
	GetInView RequestType = iota
	GetRegisteredMetricBuckets
	RegisterMetrics
	PushMetricBlob
	RegisterAggregationPlugin
	QueryMetric
	IsMetricActive
	BroadcastMessage
	AlarmTrigger
	MembershipUpdates
)

func (r RequestType) String() string {
	switch r {
	case GetInView:
		return "GetInView"
	case GetRegisteredMetricBuckets:
		return "GetRegisteredMetricBuckets"
	case RegisterMetrics:
		return "RegisterMetrics"
	case PushMetricBlob:
		return "PushMetricBlob"
	case RegisterAggregationPlugin:
		return "RegisterAggregationPlugin"
	case QueryMetric:
		return "QueryMetric"
	case IsMetricActive:
		return "IsMetricActive"
	case BroadcastMessage:
		return "BroadcastMessage"
	case AlarmTrigger:
		return "AlarmTrigger"
	case MembershipUpdates:
		return "MembershipUpdates"
	default:
		return fmt.Sprintf("%d", int(r))
	}
}

// const (

// 	//path vars
// 	ServiceNamePathVar = "serviceName"
// 	MetricNamePathVar  = "metricName"
// 	OriginNamePathVar  = "origin"

// 	GetActiveViewPath        = "/membership/view"
// 	SubscribeNodeUpdatesPath = "/membership/view/updates"
// 	GetPassiveViewPath       = "/membership/passiveView"

// 	// plugins endpoints
// 	AddPluginPath  = "/plugins"
// 	GetPluginsPath = "/plugins"

// 	// metrics_manager
// 	AddMetricsPath      = "/metrics"
// 	DeleteMetricsPath   = "/metrics/{" + ServiceNamePathVar + "}/{" + MetricNamePathVar + "}"
// 	GetMetricsPath      = "/metrics"
// 	RegisterMetricsPath = "/metrics"
// )

// // membership
// const GetPassiveViewMethod = http.MethodGet
// const GetActiveViewMethod = http.MethodGet
// const SubscribeNodeUpdatesMethod = http.MethodPost

// // plugins
// const AddPluginMethod = http.MethodPut
// const GetPluginsMethod = http.MethodGet

// // metrics
// const GetMetricsMethod = http.MethodGet
// const AddMetricsMethod = http.MethodPost
// const DeleteMetricsMethod = http.MethodDelete
// const RegisterMetricsMethod = http.MethodPut
