package requests

type StoreCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
