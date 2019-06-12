package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SignIn(ctx context.Context, email string, password string) (b0 bool, e1 error) {
	defer func() {
		l.logger.Log("method", "SignIn", "email", email, "password", password, "b0", b0, "e1", e1)
	}()
	return l.next.SignIn(ctx, email, password)
}
func (l loggingMiddleware) SignUp(ctx context.Context, email string, name string, password string) (b0 bool, e1 error) {
	defer func() {
		l.logger.Log("method", "SignUp", "email", email, "name", name, "password", password, "b0", b0, "e1", e1)
	}()
	return l.next.SignUp(ctx, email, name, password)
}
