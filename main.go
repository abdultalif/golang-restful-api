package main

import (
	"log"
	"net/http"

	"github.com/abdultalif/restful-api/app"
	"github.com/abdultalif/restful-api/controller"
	"github.com/abdultalif/restful-api/helper"
	"github.com/abdultalif/restful-api/middleware"
	"github.com/abdultalif/restful-api/repository"
	"github.com/abdultalif/restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
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
		// Handler: router,
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("🔄 Starting server...")
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}