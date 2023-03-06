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

func (h grpcHandler) PublisherAdd(ctx context.Context, r *gitspb.PublisherAddRequest) (*gitspb.PublisherAddResponse, error) {
	req := model.Publisher{
		Name: &r.GetPublisher().Name,
	}

	publisher, err := h.usecasePublisher.Add(ctx, req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.PublisherAddResponse{
		Publisher: &gitspb.Publisher{
			PublisherId: *publisher.ID,
			Name:        *publisher.Name,
		},
	}

	return response, nil
}

func (h grpcHandler) PublisherGet(ctx context.Context, r *gitspb.PublisherGetRequest) (*gitspb.PublisherGetResponse, error) {
	publisher, err := h.usecasePublisher.Get(ctx, r.GetPublisherId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.PublisherGetResponse{
		Publisher: &gitspb.Publisher{
			PublisherId: *publisher.ID,
			Name:        *publisher.Name,
		},
	}

	return response, nil
}

func (h grpcHandler) PublisherUpdate(ctx context.Context, r *gitspb.PublisherUpdateRequest) (*gitspb.PublisherUpdateResponse, error) {
	req := model.Publisher{
		Name: &r.GetPublisher().Name,
	}

	_, err := h.usecasePublisher.Update(ctx, r.GetPublisherId(), req)
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		} else if helper.IsNotFoundError(err) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.PublisherUpdateResponse{
		Result: true,
	}

	return response, nil
}

func (h grpcHandler) PublisherDelete(ctx context.Context, r *gitspb.PublisherDeleteRequest) (*gitspb.PublisherDeleteResponse, error) {
	err := h.usecasePublisher.Delete(ctx, r.GetPublisherId())
	if err != nil {
		if helper.IsParameterError(err) {
			return nil, status.Error(codes.InvalidArgument, "Incorrect request argument")
		}
		return nil, fmt.Errorf("internal server error")
	}

	response := &gitspb.PublisherDeleteResponse{
		Result: true,
	}

	return response, nil
}
