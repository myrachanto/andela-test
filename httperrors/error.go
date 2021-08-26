package httperrors
import (
	"net/http"
)
type Httperror struct {
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
	Error   string `json:"error,omitempty"`
}
func Badrequest(message string) *Httperror {
	return &Httperror{
		Message: message,
		Code: http.StatusBadRequest,
		Error: "bad request",
	}
}
func BadNotfound(message string) *Httperror {
	return &Httperror{
		Message: message,
		Code: http.StatusNotFound,
		Error: "Not found",
	}
}
