package body_types

import (
	"errors"
	"net"
	"time"

	"github.com/nm-morais/demmon-common/routes"
)

type CustomInterestSet struct {
	Query                   string
	Hosts                   []string
	OutputMetricName        string
	OutputMetricGranularity Granularity
}

type NeighborhoodInterestSet struct {
	Query                   string
	OutputMetricName        string
	OutputMetricGranularity Granularity
}

type TreeInterestSet struct {
	Query                   string
	OutputMetricName        string
	OutputMetricGranularity Granularity
}

type InstallInterestSetReply struct {
	SetId uint64
}

type Timeseries struct {
	Name   string
	Tags   map[string]string
	Points []Point
}

type Point struct {
	TS     time.Time
	Fields map[string]interface{}
}

type PointCollectionWithTagsAndName = []*PointWithTagsAndName

type PointWithTagsAndName struct {
	Point Point
	Tags  map[string]string
	Name  string
}

func NewPoint(name string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) *PointWithTagsAndName {
	return &PointWithTagsAndName{
		Name: name,
		Tags: tags,
		Point: Point{
			Fields: fields,
			TS:     timestamp,
		},
	}
}

type Granularity struct {
	Granularity time.Duration
	Count       int
}

type Peer struct {
	ID string
	IP net.IP
}

type QueryRequest struct {
	Query   string
	Timeout time.Duration
}

type InstallContinuousQueryRequest struct {
	Description       string
	FrequencySeconds  uint
	OutputMetricName  string
	OutputMetricCount int
	ExpressionTimeout time.Duration
	Expression        string
	NrRetries         int
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

type InstallContinuousQueryReply struct {
	TaskId uint64
}

type GetContinuousQueriesReply struct {
	ContinuousQueries []struct {
		TaskId    int
		NrRetries int
		CurrTry   int
		LastRan   time.Time
		Error     error
	}
}

// Request represents a request from client
type Request struct {
	ID      uint64             `json:"id"`
	Type    routes.RequestType `json:"type"`
	Message interface{}        `json:"message,omitempty"`
}

// Response is the reply message from the server
type Response struct {
	ID      uint64             `json:"id"`
	Push    bool               `json:"push"`
	Type    routes.RequestType `json:"type"`
	Error   bool               `json:"error"`
	Message interface{}        `json:"message,omitempty"`
}

func NewRequest(id uint64, reqType routes.RequestType, message interface{}) *Request {
	return &Request{
		Type:    reqType,
		ID:      id,
		Message: message,
	}
}

func NewResponse(id uint64, push bool, err error, respType routes.RequestType, message interface{}) *Response {
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
