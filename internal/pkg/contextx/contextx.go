package contextx

import "context"

type (
	requestIDKey struct {
	}
)

// WithRequestID put ID into contextual env
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// RequestID is to extract id from contextual request
func RequestID(ctx context.Context) string {
	requestID, _ := ctx.Value(requestIDKey{}).(string)
	return requestID
}
