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

// CreateBus is the resolver for the createBus field.
func (r *mutationResolver) CreateBus(ctx context.Context, input *model.CreateBusReq) (*model.CreateBusResponse, error) {
	getBus, err := r.TransportService().CreateBus(
		ctx, &order_service.CreateBusReq{
			DriverId:     input.DriverID,
			MerchantId:   input.MerchantID,
			RegionId:     input.RegionID,
			BusStatus:    input.BusStatus,
			BusName:      input.BusName,
			BusNumber:    input.BusNumber,
			BusType:      input.BusType,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			BookingCount: int32(input.BookingCount),
			PlaceCount:   int32(input.PlaceCount),
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateBus--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateBusResponse{
		ID:           getBus.Id,
		DriverID:     getBus.DriverId,
		MerchantID:   getBus.MerchantId,
		RegionID:     getBus.RegionId,
		BusStatus:    getBus.BusStatus,
		BusName:      getBus.BusName,
		BusNumber:    getBus.BusNumber,
		BusType:      getBus.BusType,
		StartedPrice: getBus.StartedPrice,
		EndPrice:     getBus.EndPrice,
		BookingCount: int(getBus.BookingCount),
		PlaceCount:   int(getBus.PlaceCount),
		CreatedAt:    getBus.CreatedAt,
		UpdatedAt:    getBus.UpdatedAt,
	}

	return resp, nil
}

// UpdateBus is the resolver for the updateBus field.
func (r *mutationResolver) UpdateBus(ctx context.Context, input *model.UpdateBusReq) (*model.UpdateBusResponse, error) {
	getBus, err := r.TransportService().UpdateBus(
		ctx, &order_service.UpdateBusReq{
			BusStatus:    input.BusStatus,
			BusName:      input.BusName,
			BusNumber:    input.BusNumber,
			BusType:      input.BusType,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			BookingCount: int32(input.BookingCount),
			PlaceCount:   int32(input.PlaceCount),
			Id:           input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateBus--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateBusResponse{
		ID:           getBus.Id,
		BusStatus:    getBus.BusStatus,
		BusName:      getBus.BusName,
		BusNumber:    getBus.BusNumber,
		BusType:      getBus.BusType,
		StartedPrice: getBus.StartedPrice,
		EndPrice:     getBus.EndPrice,
		BookingCount: int(getBus.BookingCount),
		PlaceCount:   int(getBus.PlaceCount),
		CreatedAt:    getBus.CreatedAt,
		UpdatedAt:    getBus.UpdatedAt,
	}

	return resp, nil
}

// DeleteBus is the resolver for the deleteBus field.
func (r *mutationResolver) DeleteBus(ctx context.Context, id string) (string, error) {
	_, err := r.TransportService().DeleteBus(ctx, &order_service.BusId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteBus--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListBuses is the resolver for the getListBuses field.
func (r *queryResolver) GetListBuses(ctx context.Context, input model.GetListBusesReq) (*model.GetListBusesResponse, error) {
	getBuses, err := r.TransportService().GetListBuses(
		ctx, &order_service.GetListBusesReq{Limit: int32(input.Limit), Offset: int32(input.Offset)})
	if err != nil {
		r.LoggerI.Error("!!!GetListBuses--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListBusesResponse{}
	for _, val := range getBuses.GetBuses {
		resp.GetBuses = append(resp.GetBuses, &model.GetSingleBusResponse{
			ID:           val.Id,
			DriverID:     val.DriverId,
			MerchantID:   val.MerchantId,
			RegionID:     val.RegionId,
			BusStatus:    val.BusStatus,
			BusName:      val.BusName,
			BusNumber:    val.BusNumber,
			BusType:      val.BusType,
			StartedPrice: val.StartedPrice,
			EndPrice:     val.EndPrice,
			BookingCount: int(val.BookingCount),
			PlaceCount:   int(val.PlaceCount),
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}
	resp.Count = int(getBuses.Count)

	return resp, nil
}

// GetSingleBus is the resolver for the getSingleBus field.
func (r *queryResolver) GetSingleBus(ctx context.Context, id string) (*model.GetSingleBusResponse, error) {
	getBus, err := r.TransportService().GetSingleBus(ctx, &order_service.BusId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleBus--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetSingleBusResponse{
		ID:           getBus.Id,
		DriverID:     getBus.DriverId,
		MerchantID:   getBus.MerchantId,
		RegionID:     getBus.RegionId,
		BusStatus:    getBus.BusStatus,
		BusName:      getBus.BusName,
		BusNumber:    getBus.BusNumber,
		BusType:      getBus.BusType,
		StartedPrice: getBus.StartedPrice,
		EndPrice:     getBus.EndPrice,
		BookingCount: int(getBus.BookingCount),
		PlaceCount:   int(getBus.PlaceCount),
		CreatedAt:    getBus.CreatedAt,
		UpdatedAt:    getBus.UpdatedAt,
	}

	return resp, nil
}
