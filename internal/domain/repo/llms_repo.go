package repo

import (
	"chat-email-rag-go/internal/application/dto"
	"context"
)

type LLMRepository interface {
	SendMessage(ctx context.Context, req dto.MessageReq)
	GetConversations(ctx context.Context, uuid string)
	GetConversation(ctx context.Context, req dto.ConversationDetailDTO)
	CreateConversation(ctx context.Context, req dto.CreateConversationDTO)
}
