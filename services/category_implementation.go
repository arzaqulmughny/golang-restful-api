package services

import (
	"context"
	"database/sql"
	"golang-restful-api/exceptions"
	"golang-restful-api/helpers"
	"golang-restful-api/models/domains"
	"golang-restful-api/models/requests"
	"golang-restful-api/models/resources"
	"golang-restful-api/repositories"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImplementation struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repositories.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImplementation{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImplementation) Store(ctx context.Context, request requests.StoreCategoryRequest) resources.CategoryResource {
	err := service.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := service.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	category := domains.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Store(ctx, tx, category)

	categoryResponse := resources.CategoryResource{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func (services *CategoryServiceImplementation) Update(ctx context.Context, id int, request requests.UpdateCategoryRequest) resources.CategoryResource {
	err := services.Validate.Struct(request)
	helpers.PanicIfError(err)

	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundException(err.Error()))
	}

	category.Name = request.Name
	category = services.CategoryRepository.Update(ctx, tx, id, category)
	categoryResponse := resources.CategoryResource{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func (services *CategoryServiceImplementation) Delete(ctx context.Context, id int) {
	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundException(err.Error()))
	}

	services.CategoryRepository.Delete(ctx, tx, category)
}

func (services *CategoryServiceImplementation) FindAll(ctx context.Context) []resources.CategoryResource {
	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)

	categories := services.CategoryRepository.FindAll(ctx, tx)

	var categoryResponses []resources.CategoryResource

	for _, category := range categories {
		categoryResponse := resources.CategoryResource{
			Id:   category.Id,
			Name: category.Name,
		}

		categoryResponses = append(categoryResponses, categoryResponse)
	}

	return categoryResponses
}

func (services *CategoryServiceImplementation) FindById(ctx context.Context, id int) resources.CategoryResource {
	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)

	category, err := services.CategoryRepository.FindById(ctx, tx, id)
	if err != nil {
		panic(exceptions.NewNotFoundException(err.Error()))
	}

	categoryResponse := resources.CategoryResource{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}
