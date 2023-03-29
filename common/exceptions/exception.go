package exceptions

type HTTPException[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    T      `json:"message"`
	Error      string `json:"error"`
}
