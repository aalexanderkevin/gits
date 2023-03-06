package container

import (
	"test3/config"
	"test3/repository"
	"test3/service"

	"gorm.io/gorm"
)

type Container struct {
	db     *gorm.DB
	config config.Config

	//svc
	cache service.Cache

	// repo
	bookRepo      repository.Book
	authorRepo    repository.Author
	publisherRepo repository.Publisher
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Db() *gorm.DB {
	return c.db
}

func (c *Container) SetDb(db *gorm.DB) {
	c.db = db
}

func (c *Container) Config() config.Config {
	return c.config
}

func (c *Container) SetConfig(config config.Config) {
	c.config = config
}

func (c *Container) BookRepo() repository.Book {
	return c.bookRepo
}

func (c *Container) SetBookRepo(bookRepo repository.Book) {
	c.bookRepo = bookRepo
}

func (c *Container) AuthorRepo() repository.Author {
	return c.authorRepo
}

func (c *Container) SetAuthorRepo(authorRepo repository.Author) {
	c.authorRepo = authorRepo
}

func (c *Container) PublisherRepo() repository.Publisher {
	return c.publisherRepo
}

func (c *Container) SetPublisherRepo(publisherRepo repository.Publisher) {
	c.publisherRepo = publisherRepo
}
