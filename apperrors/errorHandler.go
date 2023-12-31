package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Exzrgs/myapi/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	if !errors.As(err, &appErr) {
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParameter:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorized:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	// デバッグ用
	// fmt.Println(*appErr)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
