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

func TestBookRepository_Add(t *testing.T) {
	t.Run("ShouldInsertBook", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeBook := test.FakeBook(t, nil)

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		addedBook, err := bookRepo.Add(context.TODO(), *fakeBook)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, addedBook)
		existingBook, err := bookRepo.Get(context.TODO(), *addedBook.ID)
		require.NoError(t, err)
		require.NotNil(t, existingBook)
		require.Equal(t, addedBook.ID, existingBook.ID)
		require.Equal(t, addedBook.Name, existingBook.Name)
	})
}

func TestBookRepository_Update(t *testing.T) {
	t.Run("ShouldNotFoundError_WhenIdNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeBook := test.FakeBook(t, nil)

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		resUpdate, err := bookRepo.Update(context.Background(), "test-id", *fakeBook)

		//-- assert
		require.Error(t, err)
		require.Nil(t, resUpdate)
	})

	t.Run("ShouldUpdateBook", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeBook := test.FakeBookCreate(t, db, nil)
		fakeBook2 := model.Book{
			Name: helper.Pointer(fake.FullName()),
		}

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		resUpdate, err := bookRepo.Update(context.Background(), *fakeBook.ID, fakeBook2)
		data, err := bookRepo.Get(context.Background(), *resUpdate.ID)

		//-- assert
		require.NoError(t, err)
		require.Equal(t, *resUpdate.Name, *fakeBook2.Name)
		require.Equal(t, *resUpdate.Name, *data.Name)
		require.NotEqual(t, *fakeBook.Name, *fakeBook2.Name)
	})
}

func TestBookRepository_Get(t *testing.T) {
	t.Run("ShouldReturnError_WhenIDIsNotFound", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		res, err := bookRepo.Get(context.TODO(), "test")

		//-- assert
		require.Error(t, err)
		require.Nil(t, res)
	})

	t.Run("ShouldReturnBook", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeBook := test.FakeBookCreate(t, db, nil)

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		res, err := bookRepo.Get(context.TODO(), *fakeBook.ID)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, res)
		require.Equal(t, fakeBook.ID, res.ID)
		require.Equal(t, fakeBook.Name, res.Name)
	})
}

func TestBookRepository_Delete(t *testing.T) {
	t.Run("ShouldDeleteBook", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeBook := test.FakeBookCreate(t, db, nil)
		test.FakeBookCreate(t, db, nil)

		//-- code under test
		bookRepo := gormrepo.NewBookRepository(db)
		err := bookRepo.Delete(context.Background(), *fakeBook.ID)
		require.NoError(t, err)
		data, err := bookRepo.Get(context.Background(), *fakeBook.ID)
		require.Error(t, err)
		require.Nil(t, data)
	})
}
