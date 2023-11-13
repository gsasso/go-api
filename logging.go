package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next CustomerFetcher
}

func newLogService(next CustomerFetcher) CustomerFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchCustomer(ctx context.Context, id string) (gedi string, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"TimeSpent": time.Since(begin),
			"error":     err,
			"customer":  id,
			"Gedi":      gedi,
		}).Info("a Customer is fetched")
	}(time.Now())

	return s.next.FetchCustomer(ctx, id)
}
