package gormrepo

import (
	"test3/model"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Publisher struct {
	ID   *string
	Name *string
}

func (p Publisher) FromModel(data model.Publisher) *Publisher {
	return &Publisher{
		ID:   data.ID,
		Name: data.Name,
	}
}

func (p Publisher) ToModel() *model.Publisher {
	return &model.Publisher{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p Publisher) ToModels(publishers []Publisher) (ret []model.Publisher) {
	for _, v := range publishers {
		m := v.ToModel()
		ret = append(ret, *m)
	}
	return ret
}

func (p Publisher) GetID() *string {
	return p.ID
}

func (p Publisher) TableName() string {
	return "publishers"
}

func (p *Publisher) BeforeCreate(db *gorm.DB) error {
	if p.ID == nil {
		db.Statement.SetColumn("ID", ksuid.New().String())
	}
	return nil
}
