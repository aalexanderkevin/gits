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

func TestAuthor_Add(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenNameIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		fakeAuthor := &model.Author{}

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Add(context.Background(), *fakeAuthor)
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldAddNewAuthor", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		authorMock.On("Add", mock.Anything, *fakeAuthor).Return(fakeAuthor, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Add(context.Background(), *fakeAuthor)
		require.NoError(t, err)
		require.Equal(t, *fakeAuthor.ID, *res.ID)

		authorMock.AssertExpectations(t)
	})
}

func TestAuthor_Get(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenAuthorIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Get(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		authorMock.On("Get", mock.Anything, *fakeAuthor.ID).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Get(context.Background(), *fakeAuthor.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		authorMock.AssertExpectations(t)
	})

	t.Run("ShouldGetAuthor", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		authorMock.On("Get", mock.Anything, *fakeAuthor.ID).Return(fakeAuthor, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Get(context.Background(), *fakeAuthor.ID)
		require.NoError(t, err)
		require.Equal(t, *fakeAuthor.ID, *res.ID)

		authorMock.AssertExpectations(t)
	})
}

func TestAuthor_Update(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenAuthorIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Update(context.Background(), "", model.Author{})
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		newAuthor := model.Author{
			Name: helper.Pointer("new name"),
		}

		authorMock.On("Update", mock.Anything, *fakeAuthor.ID, newAuthor).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Update(context.Background(), *fakeAuthor.ID, newAuthor)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		authorMock.AssertExpectations(t)
	})

	t.Run("ShouldUpdateAuthor", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		newAuthor := model.Author{
			Name: helper.Pointer("new name"),
		}

		authorMock.On("Update", mock.Anything, *fakeAuthor.ID, newAuthor).Return(&newAuthor, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		res, err := uc.Update(context.Background(), *fakeAuthor.ID, newAuthor)
		require.NoError(t, err)
		require.Equal(t, *newAuthor.Name, *res.Name)

		authorMock.AssertExpectations(t)
	})
}

func TestAuthor_Delete(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenAuthorIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		err := uc.Delete(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		authorMock.On("Delete", mock.Anything, *fakeAuthor.ID).Return(errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		err := uc.Delete(context.Background(), *fakeAuthor.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())

		authorMock.AssertExpectations(t)
	})

	t.Run("ShouldDeleteAuthor", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		authorMock := &mocks.Author{}
		appContainer.SetAuthorRepo(authorMock)

		fakeAuthor := test.FakeAuthor(t, nil)

		authorMock.On("Delete", mock.Anything, *fakeAuthor.ID).Return(nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewAuthor(&appContainer)
		err := uc.Delete(context.Background(), *fakeAuthor.ID)
		require.NoError(t, err)

		authorMock.AssertExpectations(t)
	})
}
