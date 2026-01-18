package services

import (
	"context"
	"database/sql"
	"golang-restful-api/helpers"
	"golang-restful-api/models/domains"
	"golang-restful-api/models/requests"
	"golang-restful-api/models/resources"
	"golang-restful-api/repositories"
)

type CategoryServiceImplementation struct {
	CategoryRepository repositories.CategoryRepository
	DB                 *sql.DB
}

func (service *CategoryServiceImplementation) Store(ctx context.Context, request requests.StoreCategoryRequest) resources.CategoryResource {
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

func (services *CategoryServiceImplementation) Update(ctx context.Context, request requests.UpdateCategoryRequest) resources.CategoryResource {
	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	category, err := services.CategoryRepository.FindById(ctx, tx, request.Id)
	helpers.PanicIfError(err)

	category.Name = request.Name
	category = services.CategoryRepository.Update(ctx, tx, category)
	categoryResponse := resources.CategoryResource{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}

func (services *CategoryServiceImplementation) Delete(ctx context.Context, id int) {
	tx, err := services.DB.Begin()
	helpers.PanicIfError(err)

	category, err := services.CategoryRepository.FindById(ctx, tx, id)
	helpers.PanicIfError(err)

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
	helpers.PanicIfError(err)

	categoryResponse := resources.CategoryResource{
		Id:   category.Id,
		Name: category.Name,
	}

	return categoryResponse
}
