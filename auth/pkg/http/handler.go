package http

import (
	"context"
	"encoding/json"
	"errors"
	endpoint "evento_microservices/auth/pkg/endpoint"
	"net/http"
	"time"

	http1 "github.com/go-kit/kit/transport/http"
)

// makeSignUpHandler creates the handler logic
func makeSignUpHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/sign-up", http1.NewServer(endpoints.SignUpEndpoint, decodeSignUpRequest, encodeSignUpResponse, options...))
}

// decodeSignUpRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSignUpResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignUpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
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

// makeSignInHandler creates the handler logic
func makeSignInHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/sign-in", http1.NewServer(endpoints.SignInEndpoint, decodeSignInRequest, encodeSignInResponse, options...))
}

// decodeSignInRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeSignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.SignInRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeSignInResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeSignInResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}

	token := (response.(endpoint.SignInResponse)).S0
	http.SetCookie(w, &http.Cookie{
		Name:    "evento",
		Value:   token,
		Expires: time.Now().Add(10 * time.Minute),
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
