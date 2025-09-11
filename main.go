package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"nullsec45/golang-dependency-injection/app"
	"nullsec45/golang-dependency-injection/controller"
	"nullsec45/golang-dependency-injection/helper"
	"nullsec45/golang-dependency-injection/middleware"
	"nullsec45/golang-dependency-injection/repository"
	"nullsec45/golang-dependency-injection/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
