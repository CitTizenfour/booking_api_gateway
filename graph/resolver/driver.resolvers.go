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

// CreateDriver is the resolver for the createDriver field.
func (r *mutationResolver) CreateDriver(ctx context.Context, input *model.CreateDriverReq) (*model.CreateDriverResponse, error) {
	getDriver, err := r.TransportService().CreateDriver(
		ctx, &order_service.CreateDriverReq{
			PhoneNumber: input.PhoneNumber,
			Experience:  input.Experience,
			DriverName:  input.DriverName,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateDriver--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateDriverResponse{
		DriverName:  getDriver.DriverName,
		PhoneNumber: getDriver.PhoneNumber,
		Experience:  getDriver.Experience,
		ID:          getDriver.Id,
		CreatedAt:   getDriver.CreatedAt,
		UpdatedAt:   getDriver.UpdatedAt,
	}

	return resp, nil
}

// UpdateDriver is the resolver for the updateDriver field.
func (r *mutationResolver) UpdateDriver(ctx context.Context, input *model.UpdateDriverReq) (*model.UpdateDriverResponse, error) {
	getDriver, err := r.TransportService().UpdateDriver(
		ctx, &order_service.UpdateDriverReq{
			DriverName:  input.DriverName,
			PhoneNumber: input.PhoneNumber,
			Experience:  input.Experience,
			Id:          input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateDriver--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateDriverResponse{
		PhoneNumber: getDriver.PhoneNumber,
		Experience:  getDriver.Experience,
		DriverName:  getDriver.DriverName,
		ID:          getDriver.Id,
		CreatedAt:   getDriver.CreatedAt,
		UpdatedAt:   getDriver.UpdatedAt,
	}

	return resp, nil
}

// DeleteDriver is the resolver for the deleteDriver field.
func (r *mutationResolver) DeleteDriver(ctx context.Context, id string) (string, error) {
	_, err := r.TransportService().DeleteDriver(ctx, &order_service.DriverId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteDriver--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListDrivers is the resolver for the getListDrivers field.
func (r *queryResolver) GetListDrivers(ctx context.Context, input model.GetListDriversReq) (*model.GetListDriversResponse, error) {
	getDrivers, err := r.TransportService().GetListDrivers(
		ctx, &order_service.GetListDriversReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListDrivers--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListDriversResponse{}

	for _, val := range getDrivers.GetDrivers {
		resp.GetDrivers = append(resp.GetDrivers, &model.GetSingleDriverResponse{
			DriverName:  val.DriverName,
			PhoneNumber: val.PhoneNumber,
			Experience:  val.Experience,
			ID:          val.Id,
			CreatedAt:   val.CreatedAt,
			UpdatedAt:   val.UpdatedAt,
		})
	}
	resp.Count = int(getDrivers.Count)

	return resp, nil
}

// GetSingleDriver is the resolver for the getSingleDriver field.
func (r *queryResolver) GetSingleDriver(ctx context.Context, id string) (*model.GetSingleDriverResponse, error) {
	getDriver, err := r.TransportService().GetSingleDriver(ctx, &order_service.DriverId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleDriver--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetSingleDriverResponse{
		DriverName:  getDriver.DriverName,
		PhoneNumber: getDriver.PhoneNumber,
		Experience:  getDriver.Experience,
		ID:          getDriver.Id,
		CreatedAt:   getDriver.CreatedAt,
		UpdatedAt:   getDriver.UpdatedAt,
	}

	return resp, nil
}
