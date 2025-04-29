//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/fahrilhadi/golang-restful-api/app"
	"github.com/fahrilhadi/golang-restful-api/controller"
	"github.com/fahrilhadi/golang-restful-api/middleware"
	"github.com/fahrilhadi/golang-restful-api/repository"
	"github.com/fahrilhadi/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
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