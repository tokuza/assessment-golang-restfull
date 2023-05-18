package service

import (
	"assessment-golang-restful-api/exception"
	"assessment-golang-restful-api/helper"
	"assessment-golang-restful-api/model/domain"
	"assessment-golang-restful-api/model/web"
	"assessment-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(OrderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: OrderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Order := domain.Order{
		Customer_id: request.Customer_id,
		Date:        request.Date,
		Status:      request.Status,
	}

	Order = service.OrderRepository.Save(ctx, tx, Order)

	return helper.ToOrderResponse(Order)
}

func (service *OrderServiceImpl) Update(ctx context.Context, request web.OrderUpdateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Order, err := service.OrderRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	Order.Customer_id = request.Customer_id
	Order.Date = request.Date
	Order.Status = request.Status

	Order = service.OrderRepository.Update(ctx, tx, Order)

	return helper.ToOrderResponse(Order)
}

func (service *OrderServiceImpl) Delete(ctx context.Context, OrderId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Order, err := service.OrderRepository.FindById(ctx, tx, OrderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.OrderRepository.Delete(ctx, tx, Order)
}

func (service *OrderServiceImpl) FindById(ctx context.Context, OrderId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Order, err := service.OrderRepository.FindById(ctx, tx, OrderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(Order)
}

func (service *OrderServiceImpl) FindAll(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.OrderRepository.FindAll(ctx, tx)

	return helper.ToOrderResponses(categories)
}
