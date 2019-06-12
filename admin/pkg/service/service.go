package service

import "context"

// AdminService describes the service.
type AdminService interface {
	CreateEvent(
		ctx context.Context,
		UserID int,
		EventName string,
		EventDesc string,
		Date string,
	) (int, error)
	DeleteEvent(
		ctx context.Context,
		EventID int,
	) (bool, error)
}

type basicAdminService struct{}

func (b *basicAdminService) CreateEvent(ctx context.Context, UserId int, EventName string, EventDesc string, Date string) (i0 int, e1 error) {
	// TODO implement the business logic of CreateEvent
	return i0, e1
}

// NewBasicAdminService returns a naive, stateless implementation of AdminService.
func NewBasicAdminService() AdminService {
	return &basicAdminService{}
}

// New returns a AdminService with all of the expected middleware wired in.
func New(middleware []Middleware) AdminService {
	var svc AdminService = NewBasicAdminService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func (b *basicAdminService) DeleteEvent(ctx context.Context, EventId int) (b0 bool, e1 error) {
	// TODO implement the business logic of DeleteEvent
	return b0, e1
}
