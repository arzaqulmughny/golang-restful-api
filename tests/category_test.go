package tests

import (
	"database/sql"
	"encoding/json"
	"golang-restful-api/app"
	"golang-restful-api/controllers"
	"golang-restful-api/helpers"
	"golang-restful-api/middlewares"
	"golang-restful-api/repositories"
	"golang-restful-api/services"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:root1234@tcp(localhost:3306)/go_restful_api_test")

	helpers.PanicIfError(err)

	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	return db
}

func truncate(db *sql.DB) {
	db.Exec("TRUNCATE categories")
}

func setup() *middlewares.AuthMiddleware {
	db := setupDatabase()
	truncate(db)

	validate := validator.New()
	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middlewares.NewAuthMiddleware(router)
}

func TestAuthentication(t *testing.T) {
	router := setup()

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/categories", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
	assert.Equal(t, "Missing authorization header", responseBody["data"])
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
}

func TestCreateCategory(t *testing.T) {
	router := setup()

	requestBody := strings.NewReader(`{"name": "test1"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer SECRET")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "test1", responseBody["data"].(map[string]interface{})["name"])
}
