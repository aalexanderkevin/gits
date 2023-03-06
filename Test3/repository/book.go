package repository

import (
	"context"
	"test3/model"
)

type Book interface {
	Add(ctx context.Context, book model.Book) (*model.Book, error)
	Get(ctx context.Context, bookID string) (*model.Book, error)
	GetByAuthorID(ctx context.Context, author string) (*model.Book, error)
	GetByPublisherID(ctx context.Context, publisherID string) (*model.Book, error)
	Update(ctx context.Context, bookID string, notification model.Book) (*model.Book, error)
	Delete(ctx context.Context, bookID string) error
}
