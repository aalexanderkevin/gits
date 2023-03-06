//go:build integration
// +build integration

package gormrepo_test

import (
	"context"
	"test3/helper"
	"test3/helper/test"
	"test3/model"
	"test3/repository/gormrepo"
	"test3/storage"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
)

func TestPublisherRepository_Add(t *testing.T) {
	t.Run("ShouldInsertPublisher", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakePublisher := test.FakePublisher(t, nil)

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		addedPublisher, err := publisherRepo.Add(context.TODO(), *fakePublisher)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, addedPublisher)
		existingPublisher, err := publisherRepo.Get(context.TODO(), *addedPublisher.ID)
		require.NoError(t, err)
		require.NotNil(t, existingPublisher)
		require.Equal(t, addedPublisher.ID, existingPublisher.ID)
		require.Equal(t, addedPublisher.Name, existingPublisher.Name)
	})
}

func TestPublisherRepository_Update(t *testing.T) {
	t.Run("ShouldNotFoundError_WhenIdNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakePublisher := test.FakePublisher(t, nil)

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		resUpdate, err := publisherRepo.Update(context.Background(), "test-id", *fakePublisher)

		//-- assert
		require.Error(t, err)
		require.Nil(t, resUpdate)
	})

	t.Run("ShouldUpdatePublisher", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakePublisher := test.FakePublisherCreate(t, db, nil)
		fakePublisher2 := model.Publisher{
			Name: helper.Pointer(fake.FullName()),
		}

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		resUpdate, err := publisherRepo.Update(context.Background(), *fakePublisher.ID, fakePublisher2)
		data, err := publisherRepo.Get(context.Background(), *resUpdate.ID)

		//-- assert
		require.NoError(t, err)
		require.Equal(t, *resUpdate.Name, *fakePublisher2.Name)
		require.Equal(t, *resUpdate.Name, *data.Name)
		require.NotEqual(t, *fakePublisher.Name, *fakePublisher2.Name)
	})
}

func TestPublisherRepository_Get(t *testing.T) {
	t.Run("ShouldReturnError_WhenIDIsNotFound", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		res, err := publisherRepo.Get(context.TODO(), "test")

		//-- assert
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("ShouldReturnPublisher", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakePublisher := test.FakePublisherCreate(t, db, nil)

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		res, err := publisherRepo.Get(context.TODO(), *fakePublisher.ID)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, fakePublisher.ID, res.ID)
		require.Equal(t, fakePublisher.Name, res.Name)
	})
}

func TestPublisherRepository_Delete(t *testing.T) {
	t.Run("ShouldDeletePublisher", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakePublisher := test.FakePublisherCreate(t, db, nil)
		test.FakePublisherCreate(t, db, nil)

		//-- code under test
		publisherRepo := gormrepo.NewPublisherRepository(db)
		err := publisherRepo.Delete(context.Background(), *fakePublisher.ID)
		require.NoError(t, err)
		data, err := publisherRepo.Get(context.Background(), *fakePublisher.ID)
		require.Error(t, err)
		require.Nil(t, data)
	})
}
