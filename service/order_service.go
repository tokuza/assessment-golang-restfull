package service

import (
	"assessment-golang-restful-api/model/web"
	"context"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	Update(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse
	Delete(ctx context.Context, OrderId int)
	FindById(ctx context.Context, OrderId int) web.OrderResponse
	FindAll(ctx context.Context) []web.OrderResponse
}
