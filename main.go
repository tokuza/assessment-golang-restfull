package main

import (
	"assessment-golang-restful-api/app"
	"assessment-golang-restful-api/controller"
	"assessment-golang-restful-api/helper"
	"assessment-golang-restful-api/middleware"
	"assessment-golang-restful-api/repository"
	"assessment-golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	validate2 := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository, db, validate2)
	orderController := controller.NewOrderController(orderService)

	router := app.NewRouter(categoryController, orderController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
