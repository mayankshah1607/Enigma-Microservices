package endpoint

import (
	"context"
	service "evento_microservices/user/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// ShowEventRequest collects the request parameters for the ShowEvent method.
type ShowEventRequest struct {
	Id string `json:"id"`
}

// ShowEventResponse collects the response parameters for the ShowEvent method.
type ShowEventResponse struct {
	I0 interface{} `json:"i0"`
	E1 error       `json:"e1"`
}

// MakeShowEventEndpoint returns an endpoint that invokes ShowEvent on the service.
func MakeShowEventEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowEventRequest)
		i0, e1 := s.ShowEvent(ctx, req.Id)
		return ShowEventResponse{
			E1: e1,
			I0: i0,
		}, nil
	}
}

// Failed implements Failer.
func (r ShowEventResponse) Failed() error {
	return r.E1
}

// AttendEventRequest collects the request parameters for the AttendEvent method.
type AttendEventRequest struct {
	Id string `json:"id"`
}

// AttendEventResponse collects the response parameters for the AttendEvent method.
type AttendEventResponse struct {
	E0 error `json:"e0"`
}

// MakeAttendEventEndpoint returns an endpoint that invokes AttendEvent on the service.
func MakeAttendEventEndpoint(s service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AttendEventRequest)
		e0 := s.AttendEvent(ctx, req.Id)
		return AttendEventResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r AttendEventResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// ShowEvent implements Service. Primarily useful in a client.
func (e Endpoints) ShowEvent(ctx context.Context, id string) (i0 interface{}, e1 error) {
	request := ShowEventRequest{Id: id}
	response, err := e.ShowEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ShowEventResponse).I0, response.(ShowEventResponse).E1
}

// AttendEvent implements Service. Primarily useful in a client.
func (e Endpoints) AttendEvent(ctx context.Context, id string) (e0 error) {
	request := AttendEventRequest{Id: id}
	response, err := e.AttendEventEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AttendEventResponse).E0
}
