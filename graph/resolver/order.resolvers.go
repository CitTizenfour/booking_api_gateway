package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bk/bk_api_gateway/genproto/order_service"
	"bk/bk_api_gateway/graph/model"
	"bk/bk_api_gateway/pkg/logger"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.CreateOrderReq) (*model.CreateOrderResponse, error) {
	getOrder, err := r.OrderService().CreateOrder(
		ctx, &order_service.CreateOrderReq{
			UserId:     input.UserID,
			MerchantId: input.MerchantID,
		})
	if err != nil {
		r.LoggerI.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.CreateOrderResponse{
		ID:         getOrder.Id,
		UserID:     getOrder.UserId,
		MerchantID: getOrder.MerchantId,
		CreatedAt:  getOrder.CreatedAt,
		UpdatedAt:  getOrder.UpdatedAt,
	}

	return resp, nil
}

// UpdateOrder is the resolver for the updateOrder field.
func (r *mutationResolver) UpdateOrder(ctx context.Context, input *model.UpdateOrderReq) (*model.UpdateOrderResponse, error) {
	getOrder, err := r.OrderService().UpdateOrder(
		ctx, &order_service.UpdateOrderReq{
			Id:         input.ID,
			UserId:     input.UserID,
			MerchantId: input.MerchantID,
		})
	if err != nil {
		r.LoggerI.Error("!!!UpdateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	resp := &model.UpdateOrderResponse{
		ID:         getOrder.Id,
		MerchantID: getOrder.MerchantId,
		UserID:     getOrder.UserId,
		CreatedAt:  getOrder.CreatedAt,
		UpdatedAt:  getOrder.UpdatedAt,
	}

	return resp, nil
}

// DeleteOrder is the resolver for the deleteOrder field.
func (r *mutationResolver) DeleteOrder(ctx context.Context, id string) (string, error) {
	_, err := r.OrderService().DeleteOrder(ctx, &order_service.OrderId{Id: id})
	if err != nil {
		r.LoggerI.Error("!!!GetListOrders--->", logger.Error(err))
		return "Delete unsuccessful", status.Error(codes.InvalidArgument, err.Error())
	}

	return "Delete successful", nil
}

// GetListOrders is the resolver for the getListOrders field.
func (r *queryResolver) GetListOrders(ctx context.Context, input model.GetListOrdersReq) (*model.GetListOrdersResponse, error) {
	var resp model.GetListOrdersResponse
	getOrders, err := r.OrderService().GetListOrders(
		ctx, &order_service.GetListOrdersReq{
			Limit:  int32(input.Limit),
			Offset: int32(input.Offset),
		})
	if err != nil {
		r.LoggerI.Error("!!!GetListOrders--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	for _, val := range getOrders.Getorders {
		resp.Getorders = append(resp.Getorders, &model.GetSingleOrderResponse{
			ID:         val.Id,
			UserID:     val.UserId,
			MerchantID: val.MerchantId,
			CreatedAt:  val.CreatedAt,
			UpdatedAt:  val.UpdatedAt,
		})
	}
	resp.Count = int(getOrders.Count)

	return &resp, nil
}

// GetSingleOrder is the resolver for the getSingleOrder field.
func (r *queryResolver) GetSingleOrder(ctx context.Context, id string) (*model.GetSingleOrderResponse, error) {
	getOrder, err := r.OrderService().GetSingleOrder(ctx, &order_service.OrderId{Id: id})
	fmt.Println(getOrder, ``, err)
	if err != nil {
		r.LoggerI.Error("!!!GetSingleOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := &model.GetSingleOrderResponse{
		ID:         getOrder.Id,
		UserID:     getOrder.UserId,
		MerchantID: getOrder.MerchantId,
		CreatedAt:  getOrder.CreatedAt,
		UpdatedAt:  getOrder.UpdatedAt,
	}

	return resp, nil
}
