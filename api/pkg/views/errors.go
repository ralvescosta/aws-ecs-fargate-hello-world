package views

import "encoding/json"

type (
	// HTTPError
	HTTPError struct {
		StatusCode int `json:"statusCode"`
		Message    any `json:"message"`
	}
)

func (h *HTTPError) ToBuffer() []byte {
	b, _ := json.Marshal(h)
	return b
}

func UnformattedBody() *HTTPError {
	return &HTTPError{
		StatusCode: 400,
		Message:    "unformatted body",
	}
}

func BadRequest(message any) *HTTPError {
	return &HTTPError{
		StatusCode: 400,
		Message:    message,
	}
}
