package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bk/bk_api_gateway/genproto/order_service"
	"bk/bk_api_gateway/graph/model"
	"bk/bk_api_gateway/pkg/logger"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateReplyReview is the resolver for the createReplyReview field.
func (r *mutationResolver) CreateReplyReview(ctx context.Context, input *model.CreateReviewReplyReq) (*model.CreateReviewReplyResponse, error) {
	getReview, err := r.ReviewService().CreateReplyReview(
		ctx, &order_service.CreateReplyReviewsReq{
			SenderId: input.SenderID,
			ReviewId: input.ReviewID,
			Content:  input.Content,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateReplyReview--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateReviewReplyResponse{
		ID:        getReview.Id,
		SenderID:  getReview.SenderId,
		ReviewID:  getReview.ReviewId,
		Content:   getReview.Content,
		CreatedAt: getReview.CreatedAt,
		UpdatedAt: getReview.UpdatedAt,
	}

	return resp, nil
}

// UpdateReplyReview is the resolver for the updateReplyReview field.
func (r *mutationResolver) UpdateReplyReview(ctx context.Context, input *model.UpdateReviewReplyReq) (*model.UpdateReviewReplyResponse, error) {
	getReview, err := r.ReviewService().UpdateReplyReview(
		ctx, &order_service.UpdateReplyReviewReq{
			SenderId: input.SenderID,
			Content:  input.Content,
			Id:       input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateReplyReview--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.UpdateReviewReplyResponse{
		ID:        getReview.Id,
		SenderID:  getReview.SenderId,
		ReviewID:  getReview.ReviewId,
		Content:   getReview.Content,
		CreatedAt: getReview.CreatedAt,
		UpdatedAt: getReview.UpdatedAt,
	}

	return resp, nil
}

// DeleteReplyReview is the resolver for the deleteReplyReview field.
func (r *mutationResolver) DeleteReplyReview(ctx context.Context, id string) (string, error) {
	_, err := r.ReviewService().DeleteReplyReview(ctx, &order_service.ReviewId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteReplyReview--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetReplyReviews is the resolver for the GetReplyReviews field.
func (r *queryResolver) GetReplyReviews(ctx context.Context, id string) (*model.GetReplyReviewsResponse, error) {
	getReviews, err := r.ReviewService().GetReplyReviews(ctx, &order_service.ReviewId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetReplyReviews--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetReplyReviewsResponse{}
	for _, val := range getReviews.Reviewcomment {
		resp.GetReviews = append(resp.GetReviews, &model.ReviewComment{
			ID:        val.Id,
			ReviewID:  val.ReviewId,
			SenderID:  val.SenderId,
			Content:   val.Content,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	resp.Count = int(getReviews.TotalReviewComment)

	return resp, nil
}
