package logging

import (
	"context"
	"time"

	"github.com/gsasso/go-api/internal/app/service"
	"github.com/gsasso/go-api/internal/app/types"
	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next service.CustomerFetcher
}

func NewLogService(next service.CustomerFetcher) service.CustomerFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchCustomer(ctx context.Context, customReq types.CustomerRequest) (customResponse types.CustomerResponse, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value("requestID"),
			"TimeSpent": time.Since(begin),
			"error":     err,
			"customer":  customReq.Id,
			"Gedi":      customResponse.GEDI,
		}).Info("a Customer is fetched")
	}(time.Now())

	return s.next.FetchCustomer(ctx, customReq)
}
