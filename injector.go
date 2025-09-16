//go:build wireinject
// +build wireinject
package main

import (
	"github.com/google/wire"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	"nullsec45/golang-dependency-injection/app"
	"nullsec45/golang-dependency-injection/controller"
	"nullsec45/golang-dependency-injection/middleware"
	"nullsec45/golang-dependency-injection/repository"
	"nullsec45/golang-dependency-injection/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer()  *http.Server{
	wire.Build(
		app.NewDB, 
		validator.New, 
		categorySet, 
		app.NewRouter, 
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}