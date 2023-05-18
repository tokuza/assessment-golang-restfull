package app

import (
	"assessment-golang-restful-api/controller"
	"assessment-golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController, orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.GET("/api/orders", orderController.FindAll)
	router.GET("/api/orders/:orderId", orderController.FindById)
	router.POST("/api/orders", orderController.Create)
	router.PUT("/api/orders/:orderId", orderController.Update)
	router.DELETE("/api/orders/:orderId", orderController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
