package requests

type UpdateCategoryRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
