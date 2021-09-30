package render

import "fmt"

const (
	ErrCodeInternal         = 1
	ErrCodeInvalidParam     = 2
	ErrCodeUnauthorized     = 3
	ErrCodePermissionDenied = 4
	ErrCodeTimeout          = 5
)

type errResponse struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Hint    string `json:"hint,omitempty"`
}

func (e errResponse) Error() string {
	return fmt.Sprintf("%s [%d/%d]", e.Message, e.Status, e.Code)
}

func CreateError(status, code int, msg, hint string) error {
	return errResponse{
		Status:  status,
		Code:    code,
		Message: msg,
		Hint:    hint,
	}
}
