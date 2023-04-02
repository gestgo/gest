package dto

type GetListHelloQuery struct {
	Q string `json:"q" validate:"required" query:"q"`
}
