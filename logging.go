package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// func LoggingMiddleware(logger log.Logger) Middleware {
// 	return func(next OrderService) OrderService {
// 		return &loggingMiddleware{
// 			next:   next,
// 			logger: logger,
// 		}
// 	}
// }

type loggingMiddleware struct {
	logger log.Logger
	next   OrderService
}

func (mw loggingMiddleware) Create(ctx context.Context, order Order) (id string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Create",
			"CustomerID",
			order.CustomerID,
			"took", time.Since(begin),
			"err", err)
	}(time.Now())
	return mw.next.Create(ctx, order)
}

// func (mw loggingMiddleware) Count(s string) (n int) {
// 	defer func(begin time.Time) {
// 		_ = mw.logger.Log(
// 			"method", "count",
// 			"input", s,
// 			"n", n,
// 			"took", time.Since(begin),
// 		)
// 	}(time.Now())

// 	n = mw.next.Count(s)
// 	return
// }
