package observability

import (
	"context"

	core "github.com/agenticgokit/agenticgokit/internal/observability"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

// Re-export tracer configuration to make observability usable by downstream modules (e.g., CLI).
type TracerConfig = core.TracerConfig

// SetupTracer initializes the tracer provider.
func SetupTracer(ctx context.Context, cfg TracerConfig) (func(context.Context) error, error) {
	return core.SetupTracer(ctx, cfg)
}

// GetTracer returns a tracer by name.
func GetTracer(name string) trace.Tracer {
	return core.GetTracer(name)
}

// WithRunID attaches the run ID to the context.
func WithRunID(ctx context.Context, runID string) context.Context {
	return core.WithRunID(ctx, runID)
}

// RunIDFromContext retrieves the run ID from the context.
func RunIDFromContext(ctx context.Context) string {
	return core.RunIDFromContext(ctx)
}

// WithLogger attaches the logger to the context for enrichment.
func WithLogger(ctx context.Context, logger *zerolog.Logger) context.Context {
	return core.WithLogger(ctx, logger)
}

// LoggerFromContext retrieves the logger from the context if present.
func LoggerFromContext(ctx context.Context) *zerolog.Logger {
	return core.LoggerFromContext(ctx)
}

// EnrichLogger adds trace and run ID fields to the logger if available in context.
func EnrichLogger(ctx context.Context, base *zerolog.Logger) *zerolog.Logger {
	return core.EnrichLogger(ctx, base)
}
