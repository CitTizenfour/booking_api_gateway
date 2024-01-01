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

// CreateCar is the resolver for the createCar field.
func (r *mutationResolver) CreateCar(ctx context.Context, input *model.CreateCarReq) (*model.CreateCarResponse, error) {
	getCar, err := r.TransportService().CreateCar(
		ctx, &order_service.CreateCarReq{
			IsCovid:      input.IsCovid,
			IsAirCondt:   input.IsAirCondt,
			CarStatus:    input.CarStatus,
			CarNumber:    input.CarNumber,
			CarName:      input.CarName,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			PlaceCount:   int32(input.PlaceCount),
			Colour:       input.Colour,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateCar--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateCarResponse{
		IsCovid:      getCar.IsCovid,
		IsAirCondt:   getCar.IsAirCondt,
		CarStatus:    getCar.CarStatus,
		CarNumber:    getCar.CarNumber,
		CarName:      getCar.CarName,
		StartedPrice: getCar.StartedPrice,
		EndPrice:     getCar.EndPrice,
		PlaceCount:   int(getCar.PlaceCount),
		Colour:       getCar.Colour,
		ID:           getCar.Id,
		CreatedAt:    getCar.CreatedAt,
		UpdatedAt:    getCar.UpdatedAt,
	}

	return resp, nil
}

// UpdateCar is the resolver for the updateCar field.
func (r *mutationResolver) UpdateCar(ctx context.Context, input *model.UpdateCarReq) (*model.UpdateCarResponse, error) {
	getCar, err := r.TransportService().UpdateCar(
		ctx, &order_service.UpdateCarReq{
			IsCovid:      input.IsCovid,
			IsAirCondt:   input.IsAirCondt,
			CarStatus:    input.CarStatus,
			CarNumber:    input.CarNumber,
			CarName:      input.CarName,
			StartedPrice: input.StartedPrice,
			EndPrice:     input.EndPrice,
			PlaceCount:   int32(input.PlaceCount),
			Colour:       input.Colour,
			Id:           input.ID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateCar--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateCarResponse{
		IsCovid:      getCar.IsCovid,
		IsAirCondt:   getCar.IsAirCondt,
		CarStatus:    getCar.CarStatus,
		CarNumber:    getCar.CarNumber,
		CarName:      getCar.CarName,
		StartedPrice: getCar.StartedPrice,
		EndPrice:     getCar.EndPrice,
		PlaceCount:   int(getCar.PlaceCount),
		Colour:       getCar.Colour,
		ID:           getCar.Id,
		CreatedAt:    getCar.CreatedAt,
		UpdatedAt:    getCar.UpdatedAt,
	}

	return resp, nil
}

// DeleteCar is the resolver for the deleteCar field.
func (r *mutationResolver) DeleteCar(ctx context.Context, id string) (string, error) {
	_, err := r.TransportService().DeleteCar(ctx, &order_service.CarId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteCar--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListCars is the resolver for the getListCars field.
func (r *queryResolver) GetListCars(ctx context.Context, input model.GetListCarsReq) (*model.GetListCarsResponse, error) {
	getCars, err := r.TransportService().GetListCars(
		ctx, &order_service.GetListCarsReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListCars--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetListCarsResponse{}

	for _, val := range getCars.GetCars {
		resp.GetCars = append(resp.GetCars, &model.GetSingleCarResponse{
			IsCovid:      val.IsCovid,
			IsAirCondt:   val.IsAirCondt,
			CarStatus:    val.CarStatus,
			CarNumber:    val.CarNumber,
			CarName:      val.CarName,
			StartedPrice: val.StartedPrice,
			EndPrice:     val.EndPrice,
			PlaceCount:   int(val.PlaceCount),
			Colour:       val.Colour,
			ID:           val.Id,
			CreatedAt:    val.CreatedAt,
			UpdatedAt:    val.UpdatedAt,
		})
	}
	resp.Count = int(getCars.Count)

	return resp, nil
}

// GetSingleCar is the resolver for the getSingleCar field.
func (r *queryResolver) GetSingleCar(ctx context.Context, id string) (*model.GetSingleCarResponse, error) {
	getCar, err := r.TransportService().GetSingleCar(ctx, &order_service.CarId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleCar--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetSingleCarResponse{
		IsCovid:      getCar.IsCovid,
		IsAirCondt:   getCar.IsAirCondt,
		CarStatus:    getCar.CarStatus,
		CarNumber:    getCar.CarNumber,
		CarName:      getCar.CarName,
		StartedPrice: getCar.StartedPrice,
		EndPrice:     getCar.EndPrice,
		PlaceCount:   int(getCar.PlaceCount),
		Colour:       getCar.Colour,
		ID:           getCar.Id,
		CreatedAt:    getCar.CreatedAt,
		UpdatedAt:    getCar.UpdatedAt,
	}

	return resp, nil
}
