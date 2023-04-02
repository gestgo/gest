package dto

type GetHelloById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
