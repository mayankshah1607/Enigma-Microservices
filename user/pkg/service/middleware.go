package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ShowEvent(ctx context.Context, id string) (i0 interface{}, e1 error) {
	defer func() {
		l.logger.Log("method", "ShowEvent", "id", id, "i0", i0, "e1", e1)
	}()
	return l.next.ShowEvent(ctx, id)
}
func (l loggingMiddleware) AttendEvent(ctx context.Context, id string) (e0 error) {
	defer func() {
		l.logger.Log("method", "AttendEvent", "id", id, "e0", e0)
	}()
	return l.next.AttendEvent(ctx, id)
}
