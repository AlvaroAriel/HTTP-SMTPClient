package httperror

import (
	"fmt"
	"net/http"

	response "github.com/AlvaroAriel/HTTP-SMTPClient/server"
)

type APIError struct {
	StatusCode int
	Detail     string
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: status:%d, detail: %s", e.StatusCode, e.Detail)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Detail:     err.Error(),
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON"))
}

func EmptyField() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("empty field"))
}

func JSONError(w http.ResponseWriter, err error) {

	if e, ok := err.(APIError); ok {
		response.WriteJson(w, e.StatusCode, e)
		return
	}

	response.WriteJson(w, http.StatusInternalServerError, map[string]string{
		"detail": "Internal Server Error",
	})
}
