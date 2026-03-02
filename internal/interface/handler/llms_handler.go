package handler

import (
	"chat-email-rag-go/internal/application/service"

	"github.com/gin-gonic/gin"
)

type LLMHandler struct {
	llmService *service.LLMAppService
}

func NewLLMsHandler(llmService *service.LLMAppService) *LLMHandler {
	return &LLMHandler{
		llmService: llmService,
	}
}

func (l *LLMHandler) SendMessage(ctx *gin.Context) {}

func (l *LLMHandler) GetConversations(ctx *gin.Context) {}

func (l *LLMHandler) GetConversation(ctx *gin.Context) {}

func (l *LLMHandler) CreateConversations(ctx *gin.Context) {}
