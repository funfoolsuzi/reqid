package reqid

import (
	"context"

	"github.com/google/uuid"
)

// ContextKey is
type ContextKey string

const (
	// ContextKeyReqID is the context key for RequestID
	ContextKeyReqID ContextKey = "requestID"
)

// GetReqID will get reqID from a http request and return it as a string
func GetReqID(ctx context.Context) string {

	reqID := ctx.Value(ContextKeyReqID)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}

// AttachReqID will attach a brand new request ID to a http request
func AttachReqID(ctx context.Context) context.Context {

	reqID := uuid.New()

	return context.WithValue(ctx, ContextKeyReqID, reqID.String())
}
