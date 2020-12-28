package body_types

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nm-morais/demmon-common/routes"
)

// membership

func (p *Peer) String() string {
	if p == nil {
		return "<nil>"
	}

	return p.IP.String()
}

type Peer struct {
	ID string
	IP net.IP
}

type View struct {
	Children    []*Peer
	Siblings    []*Peer
	Parent      *Peer
	Grandparent *Peer
}

func (v View) String() {
	sb := &strings.Builder{}
	sb.WriteString("Children:")

	for _, c := range v.Children {
		sb.WriteString(c.String())
		sb.WriteString("|")
	}

	sb.WriteString(", Siblings:")

	for _, s := range v.Siblings {
		sb.WriteString(s.String())
		sb.WriteString("|")
	}

	sb.WriteString(", Parent:")
	sb.WriteString(v.Parent.String())

	sb.WriteString(", Grandparent:")
	sb.WriteString(v.Grandparent.String())
}

type NodeUpdateType uint

const (
	NodeUp NodeUpdateType = iota
	NodeDown
)

type NodeUpdates struct {
	Type NodeUpdateType
	Peer Peer
	View View
}

type NodeUpdateSubscriptionResponse struct {
	View View
}

// timeseries

func NewTimeseriesDTO(measurementName string, tags map[string]string, values ...ObservableDTO) TimeseriesDTO {
	ts := TimeseriesDTO{MeasurementName: measurementName, TSTags: tags, Values: values}
	return ts
}

func NewObservableDTO(fields map[string]interface{}, ts time.Time) ObservableDTO {
	return ObservableDTO{
		Fields: fields,
		TS:     ts,
	}
}

type TimeseriesDTO struct {
	MeasurementName string            `json:"name"`
	TSTags          map[string]string `json:"tags"`
	Values          []ObservableDTO   `json:"values"`
}

func (ts TimeseriesDTO) String() string {
	return fmt.Sprintf("Name: %s | Tags %+v | Values: %+v", ts.MeasurementName, ts.TSTags, ts.Values)
}

type ObservableDTO struct {
	TS     time.Time              `json:"timestamp"`
	Fields map[string]interface{} `json:"fields"`
}

func (os ObservableDTO) String() string {
	return fmt.Sprintf(" (fields:%+v, timestamp:%+v)", os.Fields, os.TS)
}

type BucketOptions struct {
	Name        string      `json:"name"`
	Granularity Granularity `json:"granularity"`
}

// type PointCollectionWithTagsAndName = []*PointWithTagsAndName

// type PointWithTagsAndName struct {
// 	Point Point
// 	Tags  map[string]string
// 	Name  string
// }

// func NewPoint(
// 	name string,
// 	tags map[string]string,
// 	fields map[string]interface{},
// 	timestamp time.Time,
// ) *PointWithTagsAndName {
// 	return &PointWithTagsAndName{
// 		Name: name,
// 		Tags: tags,
// 		Point: Point{
// 			Fields: fields,
// 			TS:     timestamp,
// 		},
// 	}
// }

type Granularity struct {
	Granularity time.Duration
	Count       int
}

// queries

type RunnableExpression struct {
	Timeout    time.Duration
	Expression string
}

type QueryRequest struct {
	Query RunnableExpression
}

type InstallContinuousQueryRequest struct {
	Description       string
	ExpressionTimeout time.Duration
	Expression        string
	NrRetries         int
	OutputBucketOpts  BucketOptions
}

type InstallContinuousQueryReply struct {
	TaskID uint64
}

type GetContinuousQueriesReply struct {
	ContinuousQueries []struct {
		TaskID    int
		NrRetries int
		CurrTry   int
		LastRan   time.Time
		Error     error
	}
}

// interest sets

type InterestSet struct {
	MaxRetries       int
	Query            RunnableExpression
	OutputBucketOpts BucketOptions
}

type CustomInterestSet struct {
	Hosts []net.IP
	IS    InterestSet
}

type NeighborhoodInterestSet struct {
	TTL int
	IS  InterestSet
}

type TreeInterestSet struct {
	Levels int
	IS     InterestSet
}

type GlobalInterestSet struct {
	IS InterestSet
}

type InstallInterestSetReply struct {
	SetID uint64
}

// broadcasts

type InstallMessageHandlerRequest struct {
	ID uint64 `json:"id"`
}

type Message struct {
	ID      uint64      `json:"id"`
	TTL     uint        `json:"ttl"`
	Content interface{} `json:"content"`
}

// Request represents a request from client

type Request struct {
	ID      uint64             `json:"id"`
	Type    routes.RequestType `json:"type"`
	Message interface{}        `json:"message,omitempty"`
}

// Response is the reply message from the server.
type Response struct {
	Message interface{}        `json:"message,omitempty"`
	ID      uint64             `json:"id"`
	Type    routes.RequestType `json:"type"`
	Push    bool               `json:"push"`
	Error   bool               `json:"error"`
	Code    int
}

func NewRequest(id uint64, reqType routes.RequestType, message interface{}) *Request {
	return &Request{
		Type:    reqType,
		ID:      id,
		Message: message,
	}
}

func NewResponse(id uint64, push bool, err error, code int, respType routes.RequestType, message interface{}) *Response {
	if err != nil {
		return &Response{
			Type:    respType,
			Push:    push,
			Error:   true,
			Code:    code,
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
		Code:    code,
	}
}

func (r *Response) GetMsgAsErr() error {
	return fmt.Errorf("status code: %d, message: %s", r.Code, errors.New(r.Message.(string)))
}
