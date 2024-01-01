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

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.CreateUserReq) (*model.CreateUserResponse, error) {
	getUser, err := r.ClienService().CreateUser(ctx, &order_service.CreateUserReq{
		Name:           input.Name,
		PaymentHistory: input.PaymentHistory,
		PhoneNumber:    input.PhoneNumber,
	})
	if err != nil {
		r.LoggerI.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateUserResponse{
		ID:             getUser.Id,
		Name:           getUser.Name,
		PaymentHistory: getUser.PaymentHistory,
		PhoneNumber:    getUser.PhoneNumber,
		CreatedAt:      getUser.CreatedAt,
		UpdatedAt:      getUser.UpdatedAt,
	}

	return resp, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.UpdateUserReq) (*model.UpdateUserResponse, error) {
	getUser, err := r.ClienService().UpdateUser(ctx, &order_service.UpdateUserReq{
		Id:             input.ID,
		Name:           input.Name,
		PaymentHistory: input.PaymentHistory,
		PhoneNumber:    input.PhoneNumber,
	})
	if err != nil {
		r.LoggerI.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.UpdateUserResponse{
		ID:             getUser.Id,
		Name:           getUser.Name,
		PaymentHistory: getUser.PaymentHistory,
		PhoneNumber:    getUser.PhoneNumber,
		CreatedAt:      getUser.CreatedAt,
		UpdatedAt:      getUser.UpdatedAt,
	}

	return resp, nil
}

// CreateMerchant is the resolver for the createMerchant field.
func (r *mutationResolver) CreateMerchant(ctx context.Context, input *model.CreateMerchantReq) (*model.CreateMerchantResponse, error) {
	getMerchant, err := r.ClienService().CreateMerchant(
		ctx, &order_service.CreateMerchantReq{
			Name:           input.Name,
			PaymentHistory: input.PaymentHistory,
			PhoneNumber:    input.PhoneNumber,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateMerchant--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.CreateMerchantResponse{
		ID:             getMerchant.Id,
		Name:           getMerchant.Name,
		PaymentHistory: getMerchant.PaymentHistory,
		PhoneNumber:    getMerchant.PhoneNumber,
		CreatedAt:      getMerchant.CreatedAt,
		UpdatedAt:      getMerchant.UpdatedAt,
	}

	return resp, nil
}

// UpdateMerchant is the resolver for the updateMerchant field.
func (r *mutationResolver) UpdateMerchant(ctx context.Context, input *model.UpdateMerchantReq) (*model.UpdateMerchantResponse, error) {
	getMerchant, err := r.ClienService().UpdateMerchant(
		ctx, &order_service.UpdateMerchantReq{
			Id:             input.ID,
			Name:           input.Name,
			PaymentHistory: input.PaymentHistory,
			PhoneNumber:    input.PhoneNumber,
		})

	if err != nil {
		r.LoggerI.Error("!!!UpdateMerchant--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.UpdateMerchantResponse{
		ID:             getMerchant.Id,
		Name:           getMerchant.Name,
		PaymentHistory: getMerchant.PaymentHistory,
		PhoneNumber:    getMerchant.PhoneNumber,
		CreatedAt:      getMerchant.CreatedAt,
		UpdatedAt:      getMerchant.UpdatedAt,
	}

	return resp, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (string, error) {
	_, err := r.ClienService().DeleteUser(ctx, &order_service.UserId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteUser--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// DeleteMerchant is the resolver for the deleteMerchant field.
func (r *mutationResolver) DeleteMerchant(ctx context.Context, id string) (string, error) {
	_, err := r.ClienService().DeleteMerchant(ctx, &order_service.MerchantId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!DeleteMerchant--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListUsers is the resolver for the getListUsers field.
func (r *queryResolver) GetListUsers(ctx context.Context, input model.GetListUsersReq) (*model.GetListUsersResponse, error) {
	getUser, err := r.ClienService().GetListUsers(ctx, &order_service.GetListUsersReq{
		Limit:  int32(input.Limit),
		Offset: int32(input.Offset),
	})
	if err != nil {
		r.LoggerI.Error("!!!GetListUsers--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetListUsersResponse{}
	for _, val := range getUser.GetUsers {
		resp.GetUsers = append(resp.GetUsers, &model.GetSingleUserResponse{
			ID:             val.Id,
			Name:           val.Name,
			PaymentHistory: val.PaymentHistory,
			PhoneNumber:    val.PhoneNumber,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}
	resp.Count = int(getUser.Count)

	return resp, nil
}

// GetSingleUser is the resolver for the getSingleUser field.
func (r *queryResolver) GetSingleUser(ctx context.Context, id string) (*model.GetSingleUserResponse, error) {
	getUser, err := r.ClienService().GetSingleUser(ctx, &order_service.UserId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetSingleUserResponse{
		ID:             getUser.Id,
		Name:           getUser.Name,
		PaymentHistory: getUser.PaymentHistory,
		PhoneNumber:    getUser.PhoneNumber,
		CreatedAt:      getUser.CreatedAt,
		UpdatedAt:      getUser.UpdatedAt,
	}

	return resp, nil
}

// GetListMerchant is the resolver for the getListMerchant field.
func (r *queryResolver) GetListMerchant(ctx context.Context, input model.GetListMerchantReq) (*model.GetListMerchantResponse, error) {
	getMerchants, err := r.ClienService().GetListMerchants(
		ctx, &order_service.GetListMerchantsReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListMerchant--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetListMerchantResponse{}
	for _, val := range getMerchants.Getmerchants {
		resp.Getmerchants = append(resp.Getmerchants, &model.GetSingleMerchantResponse{
			ID:             val.Id,
			Name:           val.Name,
			PaymentHistory: val.PaymentHistory,
			PhoneNumber:    val.PhoneNumber,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}

	resp.Count = int(getMerchants.Count)
	return resp, nil
}

// GetSingleMerchant is the resolver for the getSingleMerchant field.
func (r *queryResolver) GetSingleMerchant(ctx context.Context, id string) (*model.GetSingleMerchantResponse, error) {
	getMerchant, err := r.ClienService().GetSingleMerchant(ctx, &order_service.MerchantId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetSingleMerchant--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.GetSingleMerchantResponse{
		ID:             getMerchant.Id,
		Name:           getMerchant.Name,
		PaymentHistory: getMerchant.PaymentHistory,
		PhoneNumber:    getMerchant.PhoneNumber,
		CreatedAt:      getMerchant.CreatedAt,
		UpdatedAt:      getMerchant.UpdatedAt,
	}

	return resp, nil
}
