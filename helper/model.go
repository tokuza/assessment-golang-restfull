package helper

import (
	"assessment-golang-restful-api/model/domain"
	"assessment-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToOrderResponse(order domain.Order) web.OrderResponse {
	return web.OrderResponse{
		Id:          order.Id,
		Customer_id: order.Customer_id,
		Date:        order.Date,
		Status:      order.Status,
	}
}

func ToOrderResponses(orders []domain.Order) []web.OrderResponse {
	var orderResponses []web.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, ToOrderResponse(order))
	}
	return orderResponses
}
