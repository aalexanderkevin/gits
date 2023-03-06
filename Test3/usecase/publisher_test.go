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

func TestPublisher_Add(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenNameIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		fakePublisher := &model.Publisher{}

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Add(context.Background(), *fakePublisher)
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldAddNewPublisher", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		publisherMock.On("Add", mock.Anything, *fakePublisher).Return(fakePublisher, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Add(context.Background(), *fakePublisher)
		require.NoError(t, err)
		require.Equal(t, *fakePublisher.ID, *res.ID)

		publisherMock.AssertExpectations(t)
	})
}

func TestPublisher_Get(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenPublisherIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Get(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		publisherMock.On("Get", mock.Anything, *fakePublisher.ID).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Get(context.Background(), *fakePublisher.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		publisherMock.AssertExpectations(t)
	})

	t.Run("ShouldGetPublisher", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		publisherMock.On("Get", mock.Anything, *fakePublisher.ID).Return(fakePublisher, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Get(context.Background(), *fakePublisher.ID)
		require.NoError(t, err)
		require.Equal(t, *fakePublisher.ID, *res.ID)

		publisherMock.AssertExpectations(t)
	})
}

func TestPublisher_Update(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenPublisherIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Update(context.Background(), "", model.Publisher{})
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
		require.Nil(t, res)

	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		newPublisher := model.Publisher{
			Name: helper.Pointer("new name"),
		}

		publisherMock.On("Update", mock.Anything, *fakePublisher.ID, newPublisher).Return(nil, errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Update(context.Background(), *fakePublisher.ID, newPublisher)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())
		require.Nil(t, res)

		publisherMock.AssertExpectations(t)
	})

	t.Run("ShouldUpdatePublisher", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		newPublisher := model.Publisher{
			Name: helper.Pointer("new name"),
		}

		publisherMock.On("Update", mock.Anything, *fakePublisher.ID, newPublisher).Return(&newPublisher, nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		res, err := uc.Update(context.Background(), *fakePublisher.ID, newPublisher)
		require.NoError(t, err)
		require.Equal(t, *newPublisher.Name, *res.Name)

		publisherMock.AssertExpectations(t)
	})
}

func TestPublisher_Delete(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenPublisherIdIsMissing", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		err := uc.Delete(context.Background(), "")
		require.Error(t, err)
		require.True(t, helper.IsParameterError(err))
	})

	t.Run("ShouldError_WhenError", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		publisherMock.On("Delete", mock.Anything, *fakePublisher.ID).Return(errors.New("error")).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		err := uc.Delete(context.Background(), *fakePublisher.ID)
		require.Error(t, err)
		require.Equal(t, "error", err.Error())

		publisherMock.AssertExpectations(t)
	})

	t.Run("ShouldDeletePublisher", func(t *testing.T) {
		t.Parallel()
		// INIT
		appContainer := container.Container{}
		publisherMock := &mocks.Publisher{}
		appContainer.SetPublisherRepo(publisherMock)

		fakePublisher := test.FakePublisher(t, nil)

		publisherMock.On("Delete", mock.Anything, *fakePublisher.ID).Return(nil).Once()

		// CODE UNDER TEST
		uc := usecase.NewPublisher(&appContainer)
		err := uc.Delete(context.Background(), *fakePublisher.ID)
		require.NoError(t, err)

		publisherMock.AssertExpectations(t)
	})
}
