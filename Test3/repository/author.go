package repository

import (
	"context"
	"test3/model"
)

type Author interface {
	Add(ctx context.Context, author model.Author) (*model.Author, error)
	Get(ctx context.Context, authorID string) (*model.Author, error)
	Update(ctx context.Context, authorID string, notification model.Author) (*model.Author, error)
	Delete(ctx context.Context, authorID string) error
}
