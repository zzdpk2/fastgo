package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onexstack/fastgo/internal/pkg/contextx"
	"github.com/onexstack/fastgo/internal/pkg/known"
)

// RequestID is a middleware of Gin, which was used to inject k-v `x-request-id` for context and response from HTTP
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestID)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Save the requestID to context.Context for future call
		ctx := contextx.WithRequestID(c.Request.Context(), requestID)
		c.Request = c.Request.WithContext(ctx)

		// Put RequestID to HTTP return's HEAD. Header's key: `x-request-id`
		c.Writer.Header().Set(known.XRequestID, requestID)

		// Proceed to process the request
		c.Next()
	}
}
