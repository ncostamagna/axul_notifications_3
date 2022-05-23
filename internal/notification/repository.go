package notification

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/logger"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, notification *Notification) error
	GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB, log logger.Logger) Repository {
	return &repo{db}
}

func (r repo) Create(ctx context.Context, notification *Notification) error {
	return nil
}
func (r repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]Notification, error) {
	return nil, nil
}
