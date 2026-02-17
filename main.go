package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controllers"
	"golang-restful-api/middlewares"
	"golang-restful-api/repositories"
	"golang-restful-api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middlewares.NewAuthMiddleware(router),
	}

	server.ListenAndServe()
}
