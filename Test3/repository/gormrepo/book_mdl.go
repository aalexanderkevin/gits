package gormrepo

import (
	"test3/model"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Book struct {
	ID          *string `json:"id"`
	Name        *string `json:"no_share"`
	AuthorID    *string `json:"author_id"`
	PublisherID *string `json:"publisher_id"`
}

func (b Book) GetID() *string {
	return b.ID
}

func (b Book) TableName() string {
	return "books"
}

func (b Book) FromModel(data model.Book) *Book {
	return &Book{
		ID:          data.ID,
		Name:        data.Name,
		AuthorID:    data.AuthorID,
		PublisherID: data.PublisherID,
	}
}

func (b Book) ToModel() *model.Book {
	return &model.Book{
		ID:          b.ID,
		Name:        b.Name,
		AuthorID:    b.AuthorID,
		PublisherID: b.PublisherID,
	}
}

func (b Book) ToModels(books []Book) (ret []model.Book) {
	for _, v := range books {
		m := v.ToModel()
		ret = append(ret, *m)
	}
	return ret
}

func (b *Book) BeforeCreate(db *gorm.DB) error {
	if b.ID == nil {
		db.Statement.SetColumn("ID", ksuid.New().String())
	}
	return nil
}
