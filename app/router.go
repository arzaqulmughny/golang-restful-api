package app

import (
	"golang-restful-api/controllers"
	"golang-restful-api/exceptions"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/categories", categoryController.FindAll)
	router.POST("/categories", categoryController.Store)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exceptions.PanicHandler
	return router
}
