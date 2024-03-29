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
	// membership.
	GetInView RequestType = iota
	MembershipUpdates

	// metric buckets.
	GetRegisteredMetricBuckets
	PushMetricBlob
	InstallBucket

	// queries.
	InstallContinuousQuery
	GetContinuousQueries
	Query

	// interest sets.
	InstallCustomInterestSet
	RemoveCustomInterestSet
	UpdateCustomInterestSetHosts

	InstallNeighborhoodInterestSet

	// distributed aggregation
	InstallTreeAggregationFunction

	InstallGlobalAggregationFunction

	// broadcasts.
	BroadcastMessage
	InstallBroadcastMessageHandler

	InstallAlarm
	RemoveAlarm

	StartBabel
)

func (r RequestType) String() string {
	switch r {
	case StartBabel:
		return "StartBabel"
	case GetInView:
		return "GetInView"
	case GetRegisteredMetricBuckets:
		return "GetRegisteredMetricBuckets"
	case PushMetricBlob:
		return "PushMetricBlob"
	case InstallBucket:
		return "InstallBucket"
	case Query:
		return "QueryMetric"
	case InstallContinuousQuery:
		return "InstallContinuousQuery"
	case GetContinuousQueries:
		return "GetContinuousQueries"
	case InstallCustomInterestSet:
		return "InstallCustomInterestSet"
	case InstallGlobalAggregationFunction:
		return "InstallGlobalAggregationFunction"
	case InstallTreeAggregationFunction:
		return "InstallTreeAggregationFunction"
	case RemoveCustomInterestSet:
		return "RemoveCustomInterestSet"
	case UpdateCustomInterestSetHosts:
		return "UpdateCustomInterestSetHosts"
	case InstallNeighborhoodInterestSet:
		return "InstallNeighborhoodInterestSet"
	case BroadcastMessage:
		return "BroadcastMessage"
	case InstallBroadcastMessageHandler:
		return "InstallBroadcastMessageHandler"
	case InstallAlarm:
		return "AlarmTrigger"
	case MembershipUpdates:
		return "MembershipUpdates"
	default:
		panic(fmt.Sprintf("no String() method for requestType %d", r))
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
