package body_types

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nm-morais/demmon-common/routes"
)

var (
	ErrBadBodyType               = errors.New("bad request body type")
	ErrNonRecognizedOp           = errors.New("non-recognized operation")
	ErrCannotConnect             = errors.New("could not connect to peer")
	ErrQuerying                  = errors.New("an error occurred performing query")
	ErrCustomInterestSetNotFound = errors.New("custom interest set not found")
	ErrBucketNotFound            = errors.New("bucket not found")
)

type Peer struct {
	ID string
	IP net.IP
}

// membership

func (p *Peer) String() string {
	if p == nil {
		return "<nil>"
	}

	return p.IP.String()
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
	TaskID string
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

type CustomInterestSetErr struct {
	Err string
}

type CustomInterestSetHost struct {
	IP   net.IP
	Port int
}

type CustomInterestSet struct {
	DialRetryBackoff time.Duration
	DialTimeout      time.Duration
	Hosts            []CustomInterestSetHost
	IS               InterestSet
}

type NeighborhoodInterestSet struct {
	// stores the hop count as a tag "hop_nr" with a number corresponding to the hop in which it was registered
	StoreHopCountAsTag bool
	TTL                int
	IS                 InterestSet
}

type TreeAggregationSet struct {
	OutputBucketOpts       BucketOptions
	IntermediateBucketOpts BucketOptions
	Query                  RunnableExpression
	MergeFunction          RunnableExpression
	MaxRetries             int

	// -1 for infinite range
	Levels int

	// only relevant if UpdateOnMembershipChange == true
	MaxFrequencyUpdateOnMembershipChange time.Duration
	UpdateOnMembershipChange             bool

	StoreIntermediateValues bool
}

type GlobalAggregationFunction struct {
	Query              RunnableExpression
	MergeFunction      RunnableExpression
	DifferenceFunction RunnableExpression // assumes that first element in arguments is the minuend and the remaining args are subtrahends.
	OutputBucketOpts   BucketOptions
	MaxRetries         int

	StoreIntermediateValues bool
	IntermediateBucketOpts  BucketOptions
}

type InstallInterestSetReply struct {
	SetID string
}

type UpdateCustomInterestSetReq struct {
	SetID string
	Hosts []CustomInterestSetHost
}

type RemoveResourceRequest struct {
	ResourceID string
}

type RemoveResourceReply struct {
	ResourceID string
}

// broadcasts

type InstallMessageHandlerRequest struct {
	ID string `json:"id"`
}

type Message struct {
	ID      string      `json:"id"`
	TTL     uint        `json:"ttl"`
	Content interface{} `json:"content"`
}

// alarms

type TimeseriesFilter struct {
	MeasurementName string
	TagFilters      map[string]string
}

type InstallAlarmRequest struct {
	WatchList []TimeseriesFilter
	Query     RunnableExpression

	TriggerBackoffTime time.Duration

	MaxRetries int

	CheckPeriodicity time.Duration
	// disables only performing alarm query when the timeseries matching the watchlist are changed, saving much CPU time
	// in order to remove a timeseries when there are no values inserted, this must be set to true
	// if the number of changes to the watched timeseries would exceed the specified periodicity,
	// Demmon resorts to only executing the alarm when the suplied "CheckPeriodicity" duration has passed
	CheckPeriodic bool
}

type InstallAlarmReply struct {
	ID string `json:"id"`
}

type AlarmUpdate struct {
	ID       string `json:"id"`
	Error    bool   `json:"err"`
	Trigger  bool   `json:"trigger"`
	ErrorMsg string `json:"errMsg"`
}

// Request represents a request from client

type Request struct {
	ID      string             `json:"id"`
	Type    routes.RequestType `json:"type"`
	Message interface{}        `json:"message,omitempty"`
}

// Response is the reply message from the server.
type Response struct {
	Message interface{}        `json:"message,omitempty"`
	ID      string             `json:"id"`
	Type    routes.RequestType `json:"type"`
	Push    bool               `json:"push"`
	Error   bool               `json:"error"`
	Code    int                `json:"code"`
}

func NewRequest(id string, reqType routes.RequestType, message interface{}) *Request {
	return &Request{
		Type:    reqType,
		ID:      id,
		Message: message,
	}
}

func NewResponse(id string, push bool, err error, code int, respType routes.RequestType, message interface{}) *Response {
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
