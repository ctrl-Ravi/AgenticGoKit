package llm

import (
	"context"
	"strings"
)

// MockAdapter is a simple deterministic ModelProvider used for unit tests.
// It echoes the user prompt in a predictable format and supports basic streaming.
type MockAdapter struct {
	model string
}

// NewMockAdapter creates a new mock provider.
func NewMockAdapter(model string) *MockAdapter {
	if model == "" {
		model = "mock-model"
	}
	return &MockAdapter{model: model}
}

// Call returns a deterministic response based on the input prompt.
func (m *MockAdapter) Call(ctx context.Context, prompt Prompt) (Response, error) {
	content := "Mock response to: " + prompt.User
	resp := Response{
		Content:      content,
		Usage:        UsageStats{PromptTokens: 0, CompletionTokens: 0, TotalTokens: 0},
		FinishReason: "stop",
	}
	return resp, nil
}

// Stream returns the response content tokenized by whitespace.
func (m *MockAdapter) Stream(ctx context.Context, prompt Prompt) (<-chan Token, error) {
	ch := make(chan Token)
	go func() {
		defer close(ch)
		content := "Mock response to: " + prompt.User
		for _, w := range strings.Fields(content) {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- Token{Content: w + " "}
			}
		}
	}()
	return ch, nil
}

// Embeddings returns simple deterministic embeddings based on string lengths.
func (m *MockAdapter) Embeddings(ctx context.Context, texts []string) ([][]float64, error) {
	embeds := make([][]float64, len(texts))
	for i, t := range texts {
		// Single-dimension embedding as length; adequate for tests that don't inspect values
		embeds[i] = []float64{float64(len(t))}
	}
	return embeds, nil
}
