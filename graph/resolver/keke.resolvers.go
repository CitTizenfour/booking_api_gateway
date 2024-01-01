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

// CreateKeke is the resolver for the createKeke field.
func (r *mutationResolver) CreateKeke(ctx context.Context, input *model.CreateKekeReq) (*model.CreateKekeResponse, error) {
	getKeke, err := r.TransportService().CreateKeke(
		ctx, &order_service.CreateKekeReq{
			KekeStatus:   input.KekeStatus,
			KekeNumber:   input.KekeNumber,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			PlaceCount:   int32(input.PlaceCount),
			Colour:       input.Colour,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateKeke--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateKekeResponse{
		KekeStatus:   getKeke.KekeStatus,
		KekeNumber:   getKeke.KekeNumber,
		StartedPrice: getKeke.StartedPrice,
		EndPrice:     getKeke.EndPrice,
		PlaceCount:   int(getKeke.PlaceCount),
		Colour:       getKeke.Colour,
		ID:           getKeke.Id,
		CreatedAt:    getKeke.CreatedAt,
		UpdatedAt:    getKeke.UpdatedAt,
	}

	return resp, nil
}

// UpdateKeke is the resolver for the updateKeke field.
func (r *mutationResolver) UpdateKeke(ctx context.Context, input *model.UpdateKekeReq) (*model.UpdateKekeResponse, error) {
	getKeke, err := r.TransportService().UpdateKeke(
		ctx, &order_service.UpdateKekeReq{
			KekeStatus:   input.KekeStatus,
			KekeNumber:   input.KekeNumber,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			PlaceCount:   int32(input.PlaceCount),
			Colour:       input.Colour,
			Id:           input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateKeke--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateKekeResponse{
		KekeStatus:   getKeke.KekeStatus,
		KekeNumber:   getKeke.KekeNumber,
		StartedPrice: getKeke.StartedPrice,
		EndPrice:     getKeke.EndPrice,
		PlaceCount:   int(getKeke.PlaceCount),
		Colour:       getKeke.Colour,
		ID:           getKeke.Id,
		CreatedAt:    getKeke.CreatedAt,
		UpdatedAt:    getKeke.UpdatedAt,
	}

	return resp, nil
}

// DeleteKeke is the resolver for the deleteKeke field.
func (r *mutationResolver) DeleteKeke(ctx context.Context, id string) (string, error) {
	_, err := r.TransportService().DeleteKeke(ctx, &order_service.KekeId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteKeke--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListKekes is the resolver for the getListKekes field.
func (r *queryResolver) GetListKekes(ctx context.Context, input model.GetListKekesReq) (*model.GetListKekesResponse, error) {
	getKekes, err := r.TransportService().GetListKekes(
		ctx, &order_service.GetListKekesReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListKekes--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListKekesResponse{}

	for _, val := range getKekes.GetKekes {
		resp.GetKekes = append(resp.GetKekes, &model.GetSingleKekeResponse{
			KekeStatus:   val.KekeStatus,
			KekeNumber:   val.KekeNumber,
			StartedPrice: val.StartedPrice,
			EndPrice:     val.EndPrice,
			PlaceCount:   int(val.PlaceCount),
			Colour:       val.Colour,
			ID:           val.Id,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}
	resp.Count = int(getKekes.Count)

	return resp, nil
}

// GetSingleKeke is the resolver for the getSingleKeke field.
func (r *queryResolver) GetSingleKeke(ctx context.Context, id string) (*model.GetSingleKekeResponse, error) {
	getKeke, err := r.TransportService().GetSingleKeke(ctx, &order_service.KekeId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleKeke--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetSingleKekeResponse{
		KekeStatus:   getKeke.KekeStatus,
		KekeNumber:   getKeke.KekeNumber,
		StartedPrice: getKeke.StartedPrice,
		EndPrice:     getKeke.EndPrice,
		PlaceCount:   int(getKeke.PlaceCount),
		Colour:       getKeke.Colour,
		ID:           getKeke.Id,
		CreatedAt:    getKeke.CreatedAt,
		UpdatedAt:    getKeke.UpdatedAt,
	}

	return resp, nil
}
