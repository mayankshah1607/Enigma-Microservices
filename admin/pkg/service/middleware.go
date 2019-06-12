package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AdminService) AdminService

type loggingMiddleware struct {
	logger log.Logger
	next   AdminService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AdminService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AdminService) AdminService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateEvent(ctx context.Context, UserId int, EventName string, EventDesc string, Date string) (i0 int, e1 error) {
	defer func() {
		l.logger.Log("method", "CreateEvent", "UserId", UserId, "EventName", EventName, "EventDesc", EventDesc, "Date", Date, "i0", i0, "e1", e1)
	}()
	return l.next.CreateEvent(ctx, UserId, EventName, EventDesc, Date)
}

func (l loggingMiddleware) DeleteEvent(ctx context.Context, EventId int) (b0 bool, e1 error) {
	defer func() {
		l.logger.Log("method", "DeleteEvent", "EventId", EventId, "b0", b0, "e1", e1)
	}()
	return l.next.DeleteEvent(ctx, EventId)
}
