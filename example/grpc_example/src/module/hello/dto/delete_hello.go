package dto

type DeleteHelloById struct {
	ID string `json:"id" validate:"required" param:"id"`
}
