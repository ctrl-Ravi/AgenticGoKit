// Package providers contains internal memory provider implementations.
package providers

import (
	"github.com/agenticgokit/agenticgokit/core"
)

// NewInMemoryProvider creates a simple in-memory memory provider.
// It currently delegates to the Chromem provider with no external embedder.
func NewInMemoryProvider(config core.AgentMemoryConfig) (core.Memory, error) {
	return NewChromemProvider(config, nil)
}
