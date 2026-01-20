package requests

type UpdateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}
