package handler

import (
	"context"
	"fmt"

	"test3/helper"
	"test3/model"
	"test3/transport/grpc/gitspb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h grpcHandler) AuthorAdd(ctx context.Context, r *gitspb.AuthorAddRequest) (*gitspb.AuthorAddResponse, error) {
	req := model.Author{
		Name: &r.GetAuthor().Name,
	}

	author, err := h.usecaseAuthor.Add(ctx, req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.AuthorAddResponse{
		Author: &gitspb.Author{
			AuthorId: *author.ID,
			Name:     *author.Name,
		},
	}

	return response, nil
}

func (h grpcHandler) AuthorGet(ctx context.Context, r *gitspb.AuthorGetRequest) (*gitspb.AuthorGetResponse, error) {
	author, err := h.usecaseAuthor.Get(ctx, r.GetAuthorId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.AuthorGetResponse{
		Author: &gitspb.Author{
			AuthorId: *author.ID,
			Name:     *author.Name,
		},
	}

	return response, nil
}

func (h grpcHandler) AuthorUpdate(ctx context.Context, r *gitspb.AuthorUpdateRequest) (*gitspb.AuthorUpdateResponse, error) {
	req := model.Author{
		Name: &r.GetAuthor().Name,
	}

	_, err := h.usecaseAuthor.Update(ctx, r.GetAuthorId(), req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.AuthorUpdateResponse{
		Result: true,
	}

	return response, nil
}

func (h grpcHandler) AuthorDelete(ctx context.Context, r *gitspb.AuthorDeleteRequest) (*gitspb.AuthorDeleteResponse, error) {
	err := h.usecaseAuthor.Delete(ctx, r.GetAuthorId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.AuthorDeleteResponse{
		Result: true,
	}

	return response, nil
}
