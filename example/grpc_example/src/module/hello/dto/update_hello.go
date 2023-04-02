package dto

type UpdateHello struct {
	ID   string `json:"id" validate:"required" param:"id"`
	Name string `json:"name" validate:"required"`
}
