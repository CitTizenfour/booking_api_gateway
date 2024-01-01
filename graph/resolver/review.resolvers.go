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

// CreateReview is the resolver for the createReview field.
func (r *mutationResolver) CreateReview(ctx context.Context, input *model.CreateReviewReq) (*model.CreateReviewResponse, error) {
	getReview, err := r.ReviewService().CreateReview(
		ctx, &order_service.CreateReviewReq{
			UserId:        input.UserID,
			MerchantId:    input.MerchantID,
			Rate:          input.Rate,
			ReviewContent: input.ReviewContent,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateReview--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateReviewResponse{
		ID:            getReview.Id,
		UserID:        getReview.UserId,
		MerchantID:    getReview.MerchantId,
		Rate:          getReview.Rate,
		ReviewContent: getReview.ReviewContent,
		CreatedAt:     getReview.CreatedAt,
		UpdatedAt:     getReview.UpdatedAt,
	}

	return resp, nil
}

// UpdateReview is the resolver for the updateReview field.
func (r *mutationResolver) UpdateReview(ctx context.Context, input *model.UpdateReviewReq) (*model.UpdateReviewResponse, error) {
	getReview, err := r.ReviewService().UpdateReview(
		ctx, &order_service.UpdateReviewReq{
			Id:            input.ID,
			Rate:          input.Rate,
			ReviewContent: input.ReviewContent,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateReview--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateReviewResponse{
		ID:            getReview.Id,
		Rate:          getReview.Rate,
		ReviewContent: getReview.ReviewContent,
		CreatedAt:     getReview.CreatedAt,
		UpdatedAt:     getReview.UpdatedAt,
	}

	return resp, nil
}

// DeleteReview is the resolver for the deleteReview field.
func (r *mutationResolver) DeleteReview(ctx context.Context, id string) (string, error) {
	_, err := r.ReviewService().DeleteReview(ctx, &order_service.ReviewId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteReview--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListReviews is the resolver for the getListReviews field.
func (r *queryResolver) GetListReviews(ctx context.Context, input model.GetListReviewsReq) (*model.GetListReviewsResponse, error) {
	getReviews, err := r.ReviewService().GetListReviews(
		ctx, &order_service.GetListReviewsReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListReviews--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListReviewsResponse{}

	for _, val := range getReviews.GetReviews {
		resp.GetReviews = append(resp.GetReviews, &model.GetSingleReviewResponse{
			ID:            val.Id,
			Rate:          val.Rate,
			ReviewContent: val.ReviewContent,
			UserID:        val.UserId,
			MerchantID:    val.MerchantId,
			CreatedAt:     val.CreatedAt,
			UpdatedAt:     val.UpdatedAt,
		})
	}
	resp.Count = int(getReviews.Count)

	return resp, nil
}

// GetSingleReview is the resolver for the getSingleReview field.
func (r *queryResolver) GetSingleReview(ctx context.Context, id string) (*model.GetSingleReviewResponse, error) {
	getReview, err := r.ReviewService().GetSingleReview(ctx, &order_service.ReviewId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleReview--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetSingleReviewResponse{
		ID:            getReview.Id,
		UserID:        getReview.UserId,
		MerchantID:    getReview.MerchantId,
		Rate:          getReview.Rate,
		ReviewContent: getReview.ReviewContent,
		CreatedAt:     getReview.CreatedAt,
		UpdatedAt:     getReview.UpdatedAt,
	}

	return resp, nil
}
