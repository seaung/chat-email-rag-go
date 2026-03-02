package service

import (
	"chat-email-rag-go/internal/domain/repo"
	"context"
)

type LLMAppService struct {
	llmsRepo repo.LLMRepository
}

func NewLLMAppService(llmsRepo repo.LLMRepository) *LLMAppService {
	return &LLMAppService{
		llmsRepo: llmsRepo,
	}
}

func (l *LLMAppService) SendMessage(ctx context.Context) {}
