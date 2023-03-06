package handler

import (
	"context"
	"fmt"

	"test3/helper"
	"test3/model"
	"test3/transport/grpc/gitspb"
	"test3/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcHandler struct {
	gitspb.UnimplementedGitsServer

	usecaseBook      usecase.Book
	usecaseAuthor    usecase.Author
	usecasePublisher usecase.Publisher
}

func NewGrpcHandler(book usecase.Book, author usecase.Author, publisher usecase.Publisher) *grpcHandler {
	return &grpcHandler{
		usecaseBook:      book,
		usecaseAuthor:    author,
		usecasePublisher: publisher,
	}
}

func (h grpcHandler) BookAdd(ctx context.Context, r *gitspb.BookAddRequest) (*gitspb.BookAddResponse, error) {
	req := model.Book{
		Name:        &r.GetBook().Name,
		AuthorID:    &r.GetBook().AuthorId,
		PublisherID: &r.GetBook().PublisherId,
	}

	book, err := h.usecaseBook.Add(ctx, req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		} else if helper.IsDuplicateError(err) {
			return nil, status.Error(codes.AlreadyExists, "Data already exist")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.BookAddResponse{
		Book: &gitspb.Book{
			BookId:      *book.ID,
			Name:        *book.Name,
			AuthorId:    *book.AuthorID,
			PublisherId: *book.PublisherID,
		},
	}

	return response, nil
}

func (h grpcHandler) BookGet(ctx context.Context, r *gitspb.BookGetRequest) (*gitspb.BookGetResponse, error) {
	book, err := h.usecaseBook.Get(ctx, r.GetBookId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.BookGetResponse{
		Book: &gitspb.Book{
			BookId:      *book.ID,
			Name:        *book.Name,
			AuthorId:    *book.AuthorID,
			PublisherId: *book.PublisherID,
		},
	}

	return response, nil
}

func (h grpcHandler) BookUpdate(ctx context.Context, r *gitspb.BookUpdateRequest) (*gitspb.BookUpdateResponse, error) {
	req := model.Book{
		Name:        &r.GetBook().Name,
		AuthorID:    &r.GetBook().AuthorId,
		PublisherID: &r.GetBook().PublisherId,
	}

	_, err := h.usecaseBook.Update(ctx, r.GetBookId(), req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.BookUpdateResponse{
		Result: true,
	}

	return response, nil
}

func (h grpcHandler) BookDelete(ctx context.Context, r *gitspb.BookDeleteRequest) (*gitspb.BookDeleteResponse, error) {
	err := h.usecaseBook.Delete(ctx, r.GetBookId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.BookDeleteResponse{
		Result: true,
	}

	return response, nil
}
