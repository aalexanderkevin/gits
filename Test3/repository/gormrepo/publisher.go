package gormrepo

import (
	"context"
	"errors"
	"test3/helper"
	"test3/model"

	"gorm.io/gorm"
)

type PublisherRepository struct {
	db *gorm.DB
}

func NewPublisherRepository(db *gorm.DB) *PublisherRepository {
	return &PublisherRepository{db}
}

func (p PublisherRepository) Get(ctx context.Context, publisherID string) (*model.Publisher, error) {
	var err error

	Publisher := Publisher{
		ID: &publisherID,
	}

	err = p.db.WithContext(ctx).First(&Publisher).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError()
		}
		return nil, err
	}

	return Publisher.ToModel(), nil
}

func (p PublisherRepository) Add(ctx context.Context, publisher model.Publisher) (ret *model.Publisher, err error) {
	gormModel := Publisher{}.FromModel(publisher)

	if err = p.db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		return nil, err
	}

	return gormModel.ToModel(), nil
}

func (p PublisherRepository) Update(ctx context.Context, publisherID string, publisher model.Publisher) (resp *model.Publisher, err error) {
	_, err = p.Get(ctx, publisherID)
	if err != nil {
		return nil, err
	}

	newPublisher := Publisher{}.FromModel(publisher)

	if err = p.db.WithContext(ctx).Model(&Publisher{ID: &publisherID}).Updates(&Publisher{
		Name: newPublisher.Name,
	}).Error; err != nil {
		return nil, err
	}

	return p.Get(ctx, publisherID)
}

func (p PublisherRepository) Delete(ctx context.Context, publisherID string) error {
	return p.db.WithContext(ctx).Delete(Publisher{}, "id = ?", publisherID).Error
}
