package utils

import (
	"context"
)

// ContextKey is the content key used in paper robot context.
type ContextKey string

// Context keys.
const (
	ContextKeyEventMetadata ContextKey = "event_metadata"
)

// EventMetadata is the event metadata.
type EventMetadata struct {
	Owner       string
	Repo        string
	IssueNumber int64
}

// FromEventMetadataContext extracts current app from given context.
func FromEventMetadataContext(ctx context.Context) *EventMetadata {
	if v, ok := ctx.Value(ContextKeyEventMetadata).(*EventMetadata); ok {
		return v
	}
	return nil
}

// NewEventMetadataContext creates a new context with current app.
func NewEventMetadataContext(ctx context.Context, em *EventMetadata) context.Context {
	if ctx == nil || em == nil {
		return ctx
	}
	return context.WithValue(ctx, ContextKeyEventMetadata, em)
}
