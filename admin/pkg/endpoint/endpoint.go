package endpoint

import (
	"context"
	service "evento_microservices/admin/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateEventRequest collects the request parameters for the CreateEvent method.
type CreateEventRequest struct {
	UserId    int    `json:"user_id"`
	EventName string `json:"event_name"`
	EventDesc string `json:"event_desc"`
	Date      string `json:"date"`
}

// CreateEventResponse collects the response parameters for the CreateEvent method.
type CreateEventResponse struct {
	I0 int   `json:"i0"`
	E1 error `json:"e1"`
}

// MakeCreateEventEndpoint returns an endpoint that invokes CreateEvent on the service.
func MakeCreateEventEndpoint(s service.AdminService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEventRequest)
		i0, e1 := s.CreateEvent(ctx, req.UserId, req.EventName, req.EventDesc, req.Date)
		return CreateEventResponse{
			E1: e1,
			I0: i0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateEventResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateEvent implements Service. Primarily useful in a client.
func (e Endpoints) CreateEvent(ctx context.Context, UserId int, EventName string, EventDesc string, Date string) (i0 int, e1 error) {
	request := CreateEventRequest{
		Date:      Date,
		EventDesc: EventDesc,
		EventName: EventName,
		UserId:    UserId,
	}
	response, err := e.CreateEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateEventResponse).I0, response.(CreateEventResponse).E1
}

// DeleteEventRequest collects the request parameters for the DeleteEvent method.
type DeleteEventRequest struct {
	EventId int `json:"event_id"`
}

// DeleteEventResponse collects the response parameters for the DeleteEvent method.
type DeleteEventResponse struct {
	B0 bool  `json:"b0"`
	E1 error `json:"e1"`
}

// MakeDeleteEventEndpoint returns an endpoint that invokes DeleteEvent on the service.
func MakeDeleteEventEndpoint(s service.AdminService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEventRequest)
		b0, e1 := s.DeleteEvent(ctx, req.EventId)
		return DeleteEventResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r DeleteEventResponse) Failed() error {
	return r.E1
}

// DeleteEvent implements Service. Primarily useful in a client.
func (e Endpoints) DeleteEvent(ctx context.Context, EventId int) (b0 bool, e1 error) {
	request := DeleteEventRequest{EventId: EventId}
	response, err := e.DeleteEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DeleteEventResponse).B0, response.(DeleteEventResponse).E1
}
