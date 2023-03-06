package repository

import (
	"context"
	"test3/model"
)

type Publisher interface {
	Add(ctx context.Context, publiser model.Publisher) (*model.Publisher, error)
	Get(ctx context.Context, publisherID string) (*model.Publisher, error)
	Update(ctx context.Context, publisherID string, notification model.Publisher) (*model.Publisher, error)
	Delete(ctx context.Context, publisherID string) error
}
