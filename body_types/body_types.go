package body_types

import (
	"errors"
	"net"
	"time"

	"github.com/nm-morais/demmon-common/timeseries"
)

type Peer struct {
	ID string
	IP net.IP
}

type PluginFileBlock struct {
	Name       string
	FirstBlock bool
	FinalBlock bool
	Content    string //b64 encoded
}

type View struct {
	Children    []*Peer
	Siblings    []*Peer
	Parent      *Peer
	Grandparent *Peer
}

type NodeUpdates struct {
	Node Peer
	View View
}

type NodeUpdateSubscriptionResponse struct {
	View View
}

type MetricMetadata struct {
	Name                    string
	Service                 string
	Plugin                  string
	UnmarshalFuncSymbolName string
	MarshalFuncSymbolName   string
	Sender                  string
	Granularities           []timeseries.Granularity
}

type GlobalPropagationOptions struct {
	QueryName    string
	InputSeries  []string // this value is used to perform a prefix search and fetch all the timeseries which are used as input to the aggregation function
	FuncName     string
	Plugin       string
	Frequency    time.Duration
	OutputMetric MetricMetadata
}

type NeighbourhoodPropagationOptions struct {
	TTL          int
	OutputMetric MetricMetadata
	Frequency    time.Duration
}

// Request represents a request from client
type Request struct {
	ID      uint64      `json:"id"`
	Type    int         `json:"type"`
	Message interface{} `json:"message,omitempty"`
}

// Response is the reply message from the server
type Response struct {
	ID      uint64      `json:"id"`
	Push    bool        `json:"push"`
	Type    int         `json:"type"`
	Error   bool        `json:"error"`
	Message interface{} `json:"message,omitempty"`
}

func NewRequest(id uint64, reqType int, message interface{}) *Request {
	return &Request{
		Type:    reqType,
		ID:      id,
		Message: message,
	}
}

func NewResponse(id uint64, push bool, err error, respType int, message interface{}) *Response {
	if err != nil {
		return &Response{
			Type:    respType,
			Push:    push,
			Error:   true,
			ID:      id,
			Message: err.Error(),
		}
	}
	return &Response{
		Type:    respType,
		Push:    push,
		Error:   false,
		ID:      id,
		Message: message,
	}

}

func (r *Response) GetMsgAsErr() error {
	return errors.New(r.Message.(string))

}
