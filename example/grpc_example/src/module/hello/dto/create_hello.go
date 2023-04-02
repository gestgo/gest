package dto

type CreateHello struct {
	Name string `json:"name" validate:"required"`
}
