package gormrepo

import (
	"context"
	"errors"
	"test3/helper"
	"test3/model"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db}
}

func (b BookRepository) Get(ctx context.Context, bookID string) (*model.Book, error) {
	var err error

	bookGorm := Book{
		ID: &bookID,
	}

	err = b.db.WithContext(ctx).First(&bookGorm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError()
		}
		return nil, err
	}

	return bookGorm.ToModel(), nil
}

func (b BookRepository) GetByAuthorID(ctx context.Context, authorID string) (*model.Book, error) {
	var err error

	bookGorm := Book{
		AuthorID: &authorID,
	}

	err = b.db.WithContext(ctx).Model(&Book{}).Where("author_id = ?", authorID).First(&bookGorm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError()
		}
		return nil, err
	}

	return bookGorm.ToModel(), nil
}

func (b BookRepository) GetByPublisherID(ctx context.Context, publisherID string) (*model.Book, error) {
	var err error

	bookGorm := Book{}

	err = b.db.WithContext(ctx).Model(&Book{}).Where("publisher_id = ?", publisherID).First(&bookGorm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.NewNotFoundError()
		}
		return nil, err
	}

	return bookGorm.ToModel(), nil
}

func (b BookRepository) Add(ctx context.Context, book model.Book) (ret *model.Book, err error) {
	gormModel := Book{}.FromModel(book)

	if err = b.db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		return nil, err
	}

	return gormModel.ToModel(), nil
}

func (b BookRepository) Update(ctx context.Context, bookID string, book model.Book) (resp *model.Book, err error) {
	_, err = b.Get(ctx, bookID)
	if err != nil {
		return nil, err
	}

	newBook := Book{}.FromModel(book)

	if err = b.db.WithContext(ctx).Model(&Book{ID: &bookID}).Updates(&Book{
		Name:        newBook.Name,
		AuthorID:    newBook.AuthorID,
		PublisherID: newBook.PublisherID,
	}).Error; err != nil {
		return nil, err
	}

	return b.Get(ctx, bookID)
}

func (b BookRepository) Delete(ctx context.Context, bookID string) error {
	return b.db.WithContext(ctx).Delete(Book{}, "id = ?", bookID).Error
}
