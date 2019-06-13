package service

import (
	"context"
	models "evento_microservices/auth/pkg/db"
)

// AuthService describes the service.
type AuthService interface {
	// Add your methods here
	SignIn(ctx context.Context, email string, password string) (bool, error)
	SignUp(
		ctx context.Context,
		email string,
		name string,
		password string,
	) (bool, error)
}

type basicAuthService struct{}

func (b *basicAuthService) SignIn(ctx context.Context, email string, password string) (b0 bool, e1 error) {
	c := make(chan error)
	go models.Authorize(email, password, c)
	err := <-c
	if err != nil {
		return false, err
	}
	return true, nil
}
func (b *basicAuthService) SignUp(ctx context.Context, email string, name string, password string) (b0 bool, e1 error) {
	newObj := models.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	c := make(chan error)
	go models.CreateNew(&newObj, c)
	err := <-c
	if err != nil {
		return false, err
	}
	return true, nil
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
