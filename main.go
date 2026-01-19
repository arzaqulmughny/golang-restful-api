package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controllers"
	"golang-restful-api/repositories"
	"golang-restful-api/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db)
	categoryController := controllers.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/categories", categoryController.FindAll)
	router.POST("/categories", categoryController.Store)
	router.GET("/categories/:categoryId", categoryController.FindById)
	router.PUT("/categories/:categoryId", categoryController.Update)
	router.DELETE("/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	server.ListenAndServe()
}
