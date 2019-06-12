package service

import "context"

// UserService describes the service.
type UserService interface {
	// Add your methods here
	ShowEvent(ctx context.Context, id string) (interface{}, error)
	AttendEvent(ctx context.Context, id string) error
}

type basicUserService struct{}

func (b *basicUserService) ShowEvent(ctx context.Context, id string) (i0 interface{}, e1 error) {
	// TODO implement the business logic of ShowEvent
	return i0, e1
}
func (b *basicUserService) AttendEvent(ctx context.Context, id string) (e0 error) {
	// TODO implement the business logic of AttendEvent
	return e0
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService() UserService {
	return &basicUserService{}
}

// New returns a UserService with all of the expected middleware wired in.
func New(middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
