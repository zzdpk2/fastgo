package errorsx

import "net/http"

// Errorsx built-in standard error
var (
	// OK means succeed
	OK = &ErrorX{
		Code:    http.StatusOK,
		Message: "",
	}

	// ErrInternal represents all unknown serverside errors
	ErrInternal = &ErrorX{
		Code:    http.StatusInternalServerError,
		Reason:  "InternalError",
		Message: "Internal server error",
	}

	ErrNotFound = &ErrorX{
		Code:    http.StatusNotFound,
		Reason:  "NotFound",
		Message: "Resources not found",
	}
)
