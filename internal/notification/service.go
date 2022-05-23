package notification

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/logger"
)

type (
	Filters struct {
		ID         []string
		UserID     []string
		NotifyType []string
	}

	Service interface {
		GetAll(ctx context.Context, filters Filters, offset, limit int, pLoad string) ([]Notification, error)
	}

	service struct {
		repo   Repository
		logger logger.Logger
	}
)

func NewService(repo Repository, logger logger.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int, pLoad string) ([]Notification, error) {
	return nil, nil
}
