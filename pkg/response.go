package pkg

type Response[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}
