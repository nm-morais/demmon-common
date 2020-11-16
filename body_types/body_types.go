package body_types

import (
	"errors"
	"net"
	"time"

	"github.com/nm-morais/demmon-common/routes"
)

type PointCollection = []*Point

type Point struct {
	Name   string
	TS     int64
	Tags   map[string]string
	Fields map[string]interface{}
}

func NewPoint(name string, tags map[string]string, values map[string]interface{}, timestamp int64) *Point {
	return &Point{
		Name:   name,
		Tags:   tags,
		TS:     timestamp,
		Fields: values,
	}
}

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

type GetPluginRequest struct {
	Chunksize  int
	PluginName string
}

type NeighbourhoodPropagationOptions struct {
	TTL                  int
	PropagationFrequency time.Duration
	OutputMetricName     string
	OutputMetricTags     map[string]string
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
