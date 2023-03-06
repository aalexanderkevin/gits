package usecase

import (
	"context"
	"test3/container"
	"test3/helper"
	"test3/model"
)

type Book struct {
	app *container.Container
}

func NewBook(app *container.Container) Book {
	return Book{
		app: app,
	}
}

func (b *Book) Add(ctx context.Context, req model.Book) (*model.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// check authoe
	_, err := b.app.AuthorRepo().Get(ctx, *req.AuthorID)
	if err != nil {
		return nil, err
	}

	// check publisher
	_, err = b.app.PublisherRepo().Get(ctx, *req.PublisherID)
	if err != nil {
		return nil, err
	}

	// get publisher
	_, err = b.app.BookRepo().GetByPublisherID(ctx, *req.PublisherID)
	if !helper.IsNotFoundError(err) {
		return nil, helper.NewDuplicateError()
	}

	res, err := b.app.BookRepo().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *Book) Get(ctx context.Context, bookID string) (*model.Book, error) {
	if bookID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing bookID"))
	}

	res, err := b.app.BookRepo().Get(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *Book) Update(ctx context.Context, bookID string, req model.Book) (*model.Book, error) {
	if bookID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing bookID"))
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := b.app.BookRepo().Update(ctx, bookID, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *Book) Delete(ctx context.Context, bookID string) error {
	if bookID == "" {
		return helper.NewParameterError(helper.Pointer("missing bookID"))
	}

	err := b.app.BookRepo().Delete(ctx, bookID)
	if err != nil {
		return err
	}

	return nil
}
