package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=25,min=3" json:"name"`
}