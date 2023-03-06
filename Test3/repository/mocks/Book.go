// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "test3/model"

	mock "github.com/stretchr/testify/mock"
)

// Book is an autogenerated mock type for the Book type
type Book struct {
	mock.Mock
}

// Add provides a mock function with given fields: ctx, book
func (_m *Book) Add(ctx context.Context, book model.Book) (*model.Book, error) {
	ret := _m.Called(ctx, book)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Book) (*model.Book, error)); ok {
		return rf(ctx, book)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Book) *model.Book); ok {
		r0 = rf(ctx, book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Book) error); ok {
		r1 = rf(ctx, book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, bookID
func (_m *Book) Delete(ctx context.Context, bookID string) error {
	ret := _m.Called(ctx, bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, bookID
func (_m *Book) Get(ctx context.Context, bookID string) (*model.Book, error) {
	ret := _m.Called(ctx, bookID)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Book, error)); ok {
		return rf(ctx, bookID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Book); ok {
		r0 = rf(ctx, bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAuthorID provides a mock function with given fields: ctx, author
func (_m *Book) GetByAuthorID(ctx context.Context, author string) (*model.Book, error) {
	ret := _m.Called(ctx, author)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Book, error)); ok {
		return rf(ctx, author)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Book); ok {
		r0 = rf(ctx, author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPublisherID provides a mock function with given fields: ctx, publisherID
func (_m *Book) GetByPublisherID(ctx context.Context, publisherID string) (*model.Book, error) {
	ret := _m.Called(ctx, publisherID)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Book, error)); ok {
		return rf(ctx, publisherID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Book); ok {
		r0 = rf(ctx, publisherID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, publisherID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, bookID, notification
func (_m *Book) Update(ctx context.Context, bookID string, notification model.Book) (*model.Book, error) {
	ret := _m.Called(ctx, bookID, notification)

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Book) (*model.Book, error)); ok {
		return rf(ctx, bookID, notification)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Book) *model.Book); ok {
		r0 = rf(ctx, bookID, notification)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, model.Book) error); ok {
		r1 = rf(ctx, bookID, notification)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBook interface {
	mock.TestingT
	Cleanup(func())
}

// NewBook creates a new instance of Book. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBook(t mockConstructorTestingTNewBook) *Book {
	mock := &Book{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
