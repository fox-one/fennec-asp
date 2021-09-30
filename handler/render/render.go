package render

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

var ResponseErrorMessageAsHint bool

func init() {
	ResponseErrorMessageAsHint, _ = strconv.ParseBool(
		os.Getenv("HTTP_SHOW_ERROR_HINT"),
	)
}

func InternalServer(w http.ResponseWriter) {
	err := errResponse{
		Status:  http.StatusInternalServerError,
		Code:    ErrCodeInternal,
		Message: "internal server",
	}
	RenderError(w, err)
}

func Unauthorized(w http.ResponseWriter, hints ...string) {
	err := errResponse{
		Status:  http.StatusUnauthorized,
		Code:    ErrCodeUnauthorized,
		Message: "unauthenticated",
	}
	if len(hints) > 0 && len(hints[0]) > 0 {
		err.Hint = hints[0]
	}
	RenderError(w, err)
}

func PermissionDenied(w http.ResponseWriter) {
	err := errResponse{
		Status:  http.StatusForbidden,
		Code:    ErrCodePermissionDenied,
		Message: "permission denied",
	}
	RenderError(w, err)
}

func Timeout(w http.ResponseWriter) {
	err := errResponse{
		Status:  http.StatusBadRequest,
		Code:    ErrCodeTimeout,
		Message: "operation timeout",
	}
	RenderError(w, err)
}

func RenderError(w http.ResponseWriter, err error) {
	errResp := errResponse{}
	errResp.Status = http.StatusInternalServerError
	errResp.Message = http.StatusText(http.StatusInternalServerError)
	errResp.Code = ErrCodeInternal

	if e, ok := err.(errResponse); ok {
		errResp = e
	}
	if !ResponseErrorMessageAsHint {
		errResp.Hint = err.Error()
	}
	JSON(w, errResp, errResp.Status)
}

func Data(w http.ResponseWriter, v interface{}) {
	m := map[string]interface{}{
		"code": 0,
	}
	if v != nil {
		m["data"] = v
	}
	JSON(w, m, 200)
}

// JSON writes the json-encoded error message to the response
// with a 400 bad request status code.
func JSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.Encode(v)
}
