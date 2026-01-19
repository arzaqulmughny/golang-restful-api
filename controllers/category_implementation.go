package controllers

import (
	"encoding/json"
	"golang-restful-api/helpers"
	"golang-restful-api/models/requests"
	"golang-restful-api/models/resources"
	"golang-restful-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImplementation struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImplementation{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImplementation) Store(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	storeCategoryRequest := requests.StoreCategoryRequest{}
	err := decoder.Decode(&storeCategoryRequest)
	helpers.PanicIfError(err)

	newCategory := controller.CategoryService.Store(request.Context(), storeCategoryRequest)
	response := resources.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   newCategory,
	}

	writter.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writter)
	err = encoder.Encode(response)
	helpers.PanicIfError(err)
}

func (controller *CategoryControllerImplementation) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	updateCategoryRequest := requests.UpdateCategoryRequest{}

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	err = decoder.Decode(&updateCategoryRequest)
	helpers.PanicIfError(err)

	updatedCategory := controller.CategoryService.Update(request.Context(), id, updateCategoryRequest)
	encoder := json.NewEncoder(writter)
	response := resources.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   updatedCategory,
	}

	err = encoder.Encode(response)
	helpers.PanicIfError(err)
}

func (controller *CategoryControllerImplementation) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	response := resources.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(response)
	helpers.PanicIfError(err)
}

func (controller *CategoryControllerImplementation) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helpers.PanicIfError(err)

	category := controller.CategoryService.FindById(request.Context(), id)
	encoder := json.NewEncoder(writter)

	response := resources.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   category,
	}

	err = encoder.Encode(response)
	helpers.PanicIfError(err)
}

func (controller *CategoryControllerImplementation) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.CategoryService.FindAll(request.Context())

	response := resources.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categories,
	}

	encoder := json.NewEncoder(writter)
	err := encoder.Encode(response)
	helpers.PanicIfError(err)
}
