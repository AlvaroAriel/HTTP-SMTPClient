package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func WriteJson(w http.ResponseWriter, statusCode int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(v)
}

func JSONError(w http.ResponseWriter, err error) {

	if e, ok := err.(APIError); ok {
		WriteJson(w, e.StatusCode, e)
		return
	}

	WriteJson(w, http.StatusInternalServerError, map[string]string{
		"detail": "Internal Server Error",
	})
}
