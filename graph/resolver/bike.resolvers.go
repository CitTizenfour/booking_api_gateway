package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bk/bk_api_gateway/genproto/order_service"
	"bk/bk_api_gateway/graph/generated"
	"bk/bk_api_gateway/graph/model"
	"bk/bk_api_gateway/pkg/logger"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBike is the resolver for the createBike field.
func (r *mutationResolver) CreateBike(ctx context.Context, input *model.CreateBikeReq) (*model.CreateBikeResponse, error) {
	getBike, err := r.TransportService().CreateBike(
		ctx, &order_service.CreateBikeReq{
			BikeStatus:   input.BikeStatus,
			BikeNumber:   input.BikeNumber,
			BikeModel:    input.BikeModel,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			Colour:       input.Colour,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateBike--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateBikeResponse{
		BikeStatus:   getBike.BikeStatus,
		BikeModel:    getBike.BikeModel,
		BikeNumber:   getBike.BikeNumber,
		StartedPrice: getBike.StartedPrice,
		EndPrice:     getBike.EndPrice,
		Colour:       getBike.Colour,
		ID:           getBike.Id,
		CreatedAt:    getBike.CreatedAt,
		UpdatedAt:    getBike.UpdatedAt,
	}

	return resp, nil
}

// UpdateBike is the resolver for the updateBike field.
func (r *mutationResolver) UpdateBike(ctx context.Context, input *model.UpdateBikeReq) (*model.UpdateBikeResponse, error) {
	getBike, err := r.TransportService().UpdateBike(
		ctx, &order_service.UpdateBikeReq{
			BikeStatus:   input.BikeStatus,
			BikeNumber:   input.BikeNumber,
			BikeModel:    input.BikeModel,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			Colour:       input.Colour,
			Id:           input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateBike--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.UpdateBikeResponse{
		BikeStatus:   getBike.BikeStatus,
		BikeModel:    getBike.BikeModel,
		BikeNumber:   getBike.BikeNumber,
		StartedPrice: getBike.StartedPrice,
		EndPrice:     getBike.EndPrice,
		Colour:       getBike.Colour,
		ID:           getBike.Id,
		CreatedAt:    getBike.CreatedAt,
		UpdatedAt:    getBike.UpdatedAt,
	}

	return resp, nil
}

// DeleteBike is the resolver for the deleteBike field.
func (r *mutationResolver) DeleteBike(ctx context.Context, id string) (string, error) {
	_, err := r.TransportService().DeleteBike(ctx, &order_service.BikeId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteBike--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListBikes is the resolver for the getListBikes field.
func (r *queryResolver) GetListBikes(ctx context.Context, input model.GetListBikesReq) (*model.GetListBikesResponse, error) {
	getBikes, err := r.TransportService().GetListBikes(
		ctx, &order_service.GetListBikesReq{Limit: int32(input.Limit), Offset: int32(input.Offset)})
	if err != nil {
		r.LoggerI.Error("!!!GetListBikes--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListBikesResponse{}
	for _, val := range getBikes.GetBikes {
		resp.GetBikes = append(resp.GetBikes, &model.GetSingleBikeResponse{
			ID:           val.Id,
			BikeStatus:   val.BikeStatus,
			BikeNumber:   val.BikeNumber,
			BikeModel:    val.BikeModel,
			Colour:       val.Colour,
			StartedPrice: val.StartedPrice,
			EndPrice:     val.EndPrice,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}
	resp.Count = int(getBikes.Count)

	return resp, nil
}

// GetSingleBike is the resolver for the getSingleBike field.
func (r *queryResolver) GetSingleBike(ctx context.Context, id string) (*model.GetSingleBikeResponse, error) {
	getBike, err := r.TransportService().GetSingleBike(ctx, &order_service.BikeId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleBike--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetSingleBikeResponse{
		ID:           getBike.Id,
		BikeStatus:   getBike.BikeStatus,
		BikeModel:    getBike.BikeModel,
		BikeNumber:   getBike.BikeNumber,
		StartedPrice: getBike.StartedPrice,
		EndPrice:     getBike.EndPrice,
		Colour:       getBike.Colour,
		CreatedAt:    getBike.CreatedAt,
		UpdatedAt:    getBike.UpdatedAt,
	}

	return resp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
