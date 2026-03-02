package dto

type MessageReq struct {
	Content string `json:"content" binding:"required"`
}

type ConversationDetailDTO struct {
	UUID   string `form:"uuid"`
	ChatID string `form:"chat_id"`
}

type CreateConversationDTO struct {
	UUID    string
	Role    string
	ChatID  string
	Content string
}
