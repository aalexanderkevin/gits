package gormrepo

import (
	"test3/model"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Author struct {
	ID   *string
	Name *string
}

func (a Author) FromModel(data model.Author) *Author {
	return &Author{
		ID:   data.ID,
		Name: data.Name,
	}
}

func (a Author) ToModel() *model.Author {
	return &model.Author{
		ID:   a.ID,
		Name: a.Name,
	}
}

func (a Author) ToModels(authors []Author) (ret []model.Author) {
	for _, v := range authors {
		m := v.ToModel()
		ret = append(ret, *m)
	}
	return ret
}

func (a Author) GetID() *string {
	return a.ID
}

func (a Author) TableName() string {
	return "authors"
}

func (a *Author) BeforeCreate(db *gorm.DB) error {
	if a.ID == nil {
		db.Statement.SetColumn("ID", ksuid.New().String())
	}
	return nil
}
