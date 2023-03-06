package gormrepo

import (
	"context"
	"errors"
	"test3/helper"
	"test3/model"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db}
}

func (a AuthorRepository) Get(ctx context.Context, AuthorID string) (*model.Author, error) {
	var err error

	author := Author{
		ID: &AuthorID,
	}

	err = a.db.WithContext(ctx).First(&author).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError()
		}
		return nil, err
	}

	return author.ToModel(), nil
}

func (a AuthorRepository) Add(ctx context.Context, author model.Author) (ret *model.Author, err error) {
	gormModel := Author{}.FromModel(author)

	if err = a.db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		return nil, err
	}

	return gormModel.ToModel(), nil
}

func (a AuthorRepository) Update(ctx context.Context, authorID string, author model.Author) (resp *model.Author, err error) {
	_, err = a.Get(ctx, authorID)
	if err != nil {
		return nil, err
	}

	newAuthor := Author{}.FromModel(author)

	if err = a.db.WithContext(ctx).Model(&Author{ID: &authorID}).Updates(&Author{
		Name: newAuthor.Name,
	}).Error; err != nil {
		return nil, err
	}

	return a.Get(ctx, authorID)
}

func (a AuthorRepository) Delete(ctx context.Context, authorID string) error {
	return a.db.WithContext(ctx).Delete(Author{}, "id = ?", authorID).Error
}
