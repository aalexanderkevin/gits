package usecase_test

import (
	"context"
	"errors"
	"test3/helper"
	"test3/helper/test"
	"test3/model"
	"test3/repository/mocks"
	"test3/usecase"
	"testing"

	"test3/container"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBook_Add(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenNameIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		fakeBook := &model.Book{}

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Add(context.Background(), *fakeBook)
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldAddNewBook", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		authorMock := &mocks.Author{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetBookRepo(bookMock)
		appContainer.SetAuthorRepo(authorMock)
		appContainer.SetPublisherRepo(publisherMock)

		fakeBook := test.FakeBook(t, nil)

		bookMock.On("Add", mock.Anything, *fakeBook).Return(fakeBook, nil).Once()
		bookMock.On("GetByPublisherID", mock.Anything, *fakeBook.PublisherID).Return(nil, helper.NewNotFoundError()).Once()

		authorMock.On("Get", mock.Anything, *fakeBook.AuthorID).Return(nil, nil).Once()

		publisherMock.On("Get", mock.Anything, *fakeBook.PublisherID).Return(nil, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Add(context.Background(), *fakeBook)
		require.NoError(t, err)
		require.Equal(t, *fakeBook.ID, *res.ID)

		bookMock.AssertExpectations(t)
		authorMock.AssertExpectations(t)
		publisherMock.AssertExpectations(t)
	})
}

func TestBook_Get(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenBookIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Get(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		bookMock.On("Get", mock.Anything, *fakeBook.ID).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Get(context.Background(), *fakeBook.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		bookMock.AssertExpectations(t)
	})

	t.Run("ShouldGetBook", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		bookMock.On("Get", mock.Anything, *fakeBook.ID).Return(fakeBook, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Get(context.Background(), *fakeBook.ID)
		require.NoError(t, err)
		require.Equal(t, *fakeBook.ID, *res.ID)

		bookMock.AssertExpectations(t)
	})
}

func TestBook_Update(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenBookIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Update(context.Background(), "", model.Book{})
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		newBook := model.Book{
			Name:        helper.Pointer("new name"),
			AuthorID:    fakeBook.AuthorID,
			PublisherID: fakeBook.PublisherID,
		}

		bookMock.On("Update", mock.Anything, *fakeBook.ID, newBook).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Update(context.Background(), *fakeBook.ID, newBook)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		bookMock.AssertExpectations(t)
	})

	t.Run("ShouldUpdateBook", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		newBook := model.Book{
			Name:        helper.Pointer("new name"),
			AuthorID:    fakeBook.AuthorID,
			PublisherID: fakeBook.PublisherID,
		}

		bookMock.On("Update", mock.Anything, *fakeBook.ID, newBook).Return(&newBook, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		res, err := uc.Update(context.Background(), *fakeBook.ID, newBook)
		require.NoError(t, err)
		require.Equal(t, *newBook.Name, *res.Name)

		bookMock.AssertExpectations(t)
	})
}

func TestBook_Delete(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenBookIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		err := uc.Delete(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		bookMock.On("Delete", mock.Anything, *fakeBook.ID).Return(errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		err := uc.Delete(context.Background(), *fakeBook.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())

		bookMock.AssertExpectations(t)
	})

	t.Run("ShouldDeleteBook", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		bookMock := &mocks.Book{}
		appContainer.SetBookRepo(bookMock)

		fakeBook := test.FakeBook(t, nil)

		bookMock.On("Delete", mock.Anything, *fakeBook.ID).Return(nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewBook(&appContainer)
		err := uc.Delete(context.Background(), *fakeBook.ID)
		require.NoError(t, err)

		bookMock.AssertExpectations(t)
	})
}
