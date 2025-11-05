package core

import (
	"github.com/gin-gonic/gin"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	"net/http"
)

// ErrorResponse defines the response structure of error msg
// For unified format of error msg
type ErrorResponse struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

// WriteResponse Write response is a function for processing general response
// It will generate the successful recall or standard error response
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		// if something happen, then generate error response
		errx := errorsx.FromError(err)
		c.JSON(errx.Code, ErrorResponse{
			Reason:  errx.Reason,
			Message: errx.Message,
		})
		return
	}

	// If no error, return a successful response
	c.JSON(http.StatusOK, data)
}
