package usecase

import (
	"context"
	"test3/container"
	"test3/helper"
	"test3/model"
)

type Publisher struct {
	app *container.Container
}

func NewPublisher(app *container.Container) Publisher {
	return Publisher{
		app: app,
	}
}

func (p *Publisher) Add(ctx context.Context, req model.Publisher) (*model.Publisher, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := p.app.PublisherRepo().Add(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Publisher) Get(ctx context.Context, publisherID string) (*model.Publisher, error) {
	if publisherID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing publisherID"))
	}

	res, err := p.app.PublisherRepo().Get(ctx, publisherID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Publisher) Update(ctx context.Context, publisherID string, req model.Publisher) (*model.Publisher, error) {
	if publisherID == "" {
		return nil, helper.NewParameterError(helper.Pointer("missing publisherID"))
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := p.app.PublisherRepo().Update(ctx, publisherID, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *Publisher) Delete(ctx context.Context, publisherID string) error {
	if publisherID == "" {
		return helper.NewParameterError(helper.Pointer("missing publisherID"))
	}

	err := p.app.PublisherRepo().Delete(ctx, publisherID)
	if err != nil {
		return err
	}

	return nil
}
