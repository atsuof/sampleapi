package apperrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusInternalServerError

	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest

	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	if encErr := json.NewEncoder(w).Encode(appErr); encErr != nil {
		http.Error(w, fmt.Sprintf("unknown error occured: %v", encErr), http.StatusInternalServerError)
	}
}
