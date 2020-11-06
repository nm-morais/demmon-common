package routes

const (
	Dial = "/dial"
)

const (
	GetInView = iota

	GetRegisteredMetrics
	GetRegisteredPlugins

	RegisterMetrics
	PushMetricBlob

	QueryMetric
	IsMetricActive

	// plugins
	AddPlugin
	BroadcastMessage

	AlarmTrigger
	MembershipUpdates
)

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
