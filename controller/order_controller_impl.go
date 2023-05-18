package controller

import (
	"assessment-golang-restful-api/helper"
	"assessment-golang-restful-api/model/web"
	"assessment-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(OrderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: OrderService,
	}
}

func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(request.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderUpdateRequest := web.OrderUpdateRequest{}
	helper.ReadFromRequestBody(request, &orderUpdateRequest)

	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderUpdateRequest.Id = id

	orderResponse := controller.OrderService.Update(request.Context(), orderUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	controller.OrderService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderId := params.ByName("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderResponses := controller.OrderService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
