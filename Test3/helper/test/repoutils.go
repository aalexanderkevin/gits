package test

import (
	"context"
	"test3/helper"
	"test3/model"
	"test3/repository/gormrepo"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func FakeBook(t *testing.T, cb func(m model.Book) model.Book) *model.Book {
	t.Helper()
	book := model.Book{
		ID:          helper.Pointer(fake.CharactersN(5)),
		Name:        helper.Pointer(fake.FullName()),
		AuthorID:    helper.Pointer(fake.CharactersN(4)),
		PublisherID: helper.Pointer(fake.CharactersN(5)),
	}

	if cb != nil {
		book = cb(book)
	}

	return &book
}

func FakeBookCreate(t *testing.T, db *gorm.DB, cb func(m model.Book) model.Book) *model.Book {
	t.Helper()

	fakeBook := FakeBook(t, cb)

	tenantUserRepo := gormrepo.NewBookRepository(db)
	_, err := tenantUserRepo.Add(context.Background(), *fakeBook)
	require.NoError(t, err)

	return fakeBook
}

func FakeAuthor(t *testing.T, cb func(m model.Author) model.Author) *model.Author {
	t.Helper()
	author := model.Author{
		ID:   helper.Pointer(fake.CharactersN(5)),
		Name: helper.Pointer(fake.FullName()),
	}

	if cb != nil {
		author = cb(author)
	}
	return &author
}

func FakeAuthorCreate(t *testing.T, db *gorm.DB, cb func(m model.Author) model.Author) *model.Author {
	t.Helper()

	fakeAuthor := FakeAuthor(t, cb)

	tenantUserRepo := gormrepo.NewAuthorRepository(db)
	_, err := tenantUserRepo.Add(context.Background(), *fakeAuthor)
	require.NoError(t, err)

	return fakeAuthor
}

func FakePublisher(t *testing.T, cb func(m model.Publisher) model.Publisher) *model.Publisher {
	t.Helper()
	publisher := model.Publisher{
		ID:   helper.Pointer(fake.CharactersN(5)),
		Name: helper.Pointer(fake.FullName()),
	}

	if cb != nil {
		publisher = cb(publisher)
	}
	return &publisher
}

func FakePublisherCreate(t *testing.T, db *gorm.DB, cb func(m model.Publisher) model.Publisher) *model.Publisher {
	t.Helper()

	fakePublisher := FakePublisher(t, cb)

	tenantUserRepo := gormrepo.NewPublisherRepository(db)
	_, err := tenantUserRepo.Add(context.Background(), *fakePublisher)
	require.NoError(t, err)

	return fakePublisher
}
