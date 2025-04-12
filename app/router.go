package app

import (
	"github.com/abdultalif/restful-api/controller"
	"github.com/abdultalif/restful-api/error"
	"github.com/julienschmidt/httprouter"
)


func NewRouter(categoryController controller. CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/categories", categoryController.FindAll)		
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)		
	router.POST("/api/v1/categories", categoryController.Create)		
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)		
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)		

	router.PanicHandler = error.ErrorHandler
	return router
}