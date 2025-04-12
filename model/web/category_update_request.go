package web

type CategoryUpdateRequest struct {
	Id   int
	Name string `validate:"required,max=255,min=3" json:"name"`
}
