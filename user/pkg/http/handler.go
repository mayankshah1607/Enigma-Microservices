package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "evento_microservices/user/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
	"net/http"
)

// makeShowEventHandler creates the handler logic
func makeShowEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/show-event", http1.NewServer(endpoints.ShowEventEndpoint, decodeShowEventRequest, encodeShowEventResponse, options...))
}

// decodeShowEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeShowEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.ShowEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeShowEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeShowEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeAttendEventHandler creates the handler logic
func makeAttendEventHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/attend-event", http1.NewServer(endpoints.AttendEventEndpoint, decodeAttendEventRequest, encodeAttendEventResponse, options...))
}

// decodeAttendEventRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeAttendEventRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.AttendEventRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeAttendEventResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeAttendEventResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
