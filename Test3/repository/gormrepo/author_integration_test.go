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

func TestAuthorRepository_Add(t *testing.T) {
	t.Run("ShouldInsertAuthor", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeAuthor := test.FakeAuthor(t, nil)

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		addedAuthor, err := authorRepo.Add(context.TODO(), *fakeAuthor)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, addedAuthor)
		existingAuthor, err := authorRepo.Get(context.TODO(), *addedAuthor.ID)
		require.NoError(t, err)
		require.NotNil(t, existingAuthor)
		require.Equal(t, addedAuthor.ID, existingAuthor.ID)
		require.Equal(t, addedAuthor.Name, existingAuthor.Name)
	})
}

func TestAuthorRepository_Update(t *testing.T) {
	t.Run("ShouldNotFoundError_WhenIdNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeAuthor := test.FakeAuthor(t, nil)

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		resUpdate, err := authorRepo.Update(context.Background(), "test-id", *fakeAuthor)

		//-- assert
		require.Error(t, err)
		require.Nil(t, resUpdate)
	})

	t.Run("ShouldUpdateAuthor", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeAuthor := test.FakeAuthorCreate(t, db, nil)
		fakeAuthor2 := model.Author{
			Name: helper.Pointer(fake.FullName()),
		}

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		resUpdate, err := authorRepo.Update(context.Background(), *fakeAuthor.ID, fakeAuthor2)
		data, err := authorRepo.Get(context.Background(), *resUpdate.ID)

		//-- assert
		require.NoError(t, err)
		require.Equal(t, *resUpdate.Name, *fakeAuthor2.Name)
		require.Equal(t, *resUpdate.Name, *data.Name)
		require.NotEqual(t, *fakeAuthor.Name, *fakeAuthor2.Name)
	})
}

func TestAuthorRepository_Get(t *testing.T) {
	t.Run("ShouldReturnError_WhenIDIsNotFound", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		res, err := authorRepo.Get(context.TODO(), "test")

		//-- assert
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("ShouldReturnAuthor", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeAuthor := test.FakeAuthorCreate(t, db, nil)

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		res, err := authorRepo.Get(context.TODO(), *fakeAuthor.ID)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, fakeAuthor.ID, res.ID)
		require.Equal(t, fakeAuthor.Name, res.Name)
	})
}

func TestAuthorRepository_Delete(t *testing.T) {
	t.Run("ShouldDeleteAuthor", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeAuthor := test.FakeAuthorCreate(t, db, nil)
		test.FakeAuthorCreate(t, db, nil)

		//-- code under test
		authorRepo := gormrepo.NewAuthorRepository(db)
		err := authorRepo.Delete(context.Background(), *fakeAuthor.ID)
		require.NoError(t, err)
		data, err := authorRepo.Get(context.Background(), *fakeAuthor.ID)
		require.Error(t, err)
		require.Nil(t, data)
	})
}
