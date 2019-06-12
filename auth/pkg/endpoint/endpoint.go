package endpoint

import (
	"context"
	service "evento_microservices/auth/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// SignInRequest collects the request parameters for the SignIn method.
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignInResponse collects the response parameters for the SignIn method.
type SignInResponse struct {
	B0 bool  `json:"b0"`
	E1 error `json:"e1"`
}

// MakeSignInEndpoint returns an endpoint that invokes SignIn on the service.
func MakeSignInEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignInRequest)
		b0, e1 := s.SignIn(ctx, req.Email, req.Password)
		return SignInResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r SignInResponse) Failed() error {
	return r.E1
}

// SignUpRequest collects the request parameters for the SignUp method.
type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// SignUpResponse collects the response parameters for the SignUp method.
type SignUpResponse struct {
	B0 bool  `json:"b0"`
	E1 error `json:"e1"`
}

// MakeSignUpEndpoint returns an endpoint that invokes SignUp on the service.
func MakeSignUpEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		b0, e1 := s.SignUp(ctx, req.Email, req.Name, req.Password)
		return SignUpResponse{
			B0: b0,
			E1: e1,
		}, nil
	}
}

// Failed implements Failer.
func (r SignUpResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SignIn implements Service. Primarily useful in a client.
func (e Endpoints) SignIn(ctx context.Context, email string, password string) (b0 bool, e1 error) {
	request := SignInRequest{
		Email:    email,
		Password: password,
	}
	response, err := e.SignInEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignInResponse).B0, response.(SignInResponse).E1
}

// SignUp implements Service. Primarily useful in a client.
func (e Endpoints) SignUp(ctx context.Context, email string, name string, password string) (b0 bool, e1 error) {
	request := SignUpRequest{
		Email:    email,
		Name:     name,
		Password: password,
	}
	response, err := e.SignUpEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SignUpResponse).B0, response.(SignUpResponse).E1
}
