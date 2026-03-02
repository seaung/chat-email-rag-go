package agent

import (
	"chat-email-rag-go/pkg/documents"

	"github.com/tmc/langchaingo/llms"
)

type Agent struct {
	llm       llms.Model
	docLoader *documents.DocumentLoader
}

func NewAgent() (*Agent, error) {
	return &Agent{}, nil
}
