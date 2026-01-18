package services

import (
	"context"
	"golang-restful-api/models/requests"
	"golang-restful-api/models/resources"
)

type CategoryService interface {
	Store(ctx context.Context, request requests.StoreCategoryRequest) resources.CategoryResource
	Update(ctx context.Context, request requests.UpdateCategoryRequest) resources.CategoryResource
	Delete(ctx context.Context, id int)
	FindAll(ctx context.Context) []resources.CategoryResource
	FindById(ctx context.Context, id int) resources.CategoryResource
}
