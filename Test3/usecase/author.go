package usecase

import (
	"context"
	"test3/container"
	"test3/helper"
	"test3/model"
)

type Author struct {
	app *container.Container
}

func NewAuthor(app *container.Container) Author {
	return Author{
		app: app,
	}
}

func (a *Author) Add(ctx context.Context, req model.Author) (*model.Author, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := a.app.AuthorRepo().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Author) Get(ctx context.Context, authorID string) (*model.Author, error) {
	if authorID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing authorID"))
	}

	res, err := a.app.AuthorRepo().Get(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Author) Update(ctx context.Context, authorID string, req model.Author) (*model.Author, error) {
	if authorID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing authorID"))
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := a.app.AuthorRepo().Update(ctx, authorID, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *Author) Delete(ctx context.Context, authorID string) error {
	if authorID == "" {
		return helper.NewParameterError(helper.Pointer("missing authorID"))
	}

	err := a.app.AuthorRepo().Delete(ctx, authorID)
	if err != nil {
		return err
	}

	return nil
}
