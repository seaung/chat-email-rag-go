package persistence

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Name         string    `json:"name"`
	Avatar       string    `json:"avatar"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// 关联
	Conversations []Conversation `json:"conversations,omitempty" gorm:"foreignKey:UserID"`
	Files         []File         `json:"files,omitempty" gorm:"foreignKey:UserID"`
}

// Conversation 对话模型
type Conversation struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Title     string    `json:"title" gorm:"default:'新对话'"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联
	User     User      `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Messages []Message `json:"messages,omitempty" gorm:"foreignKey:ConversationID"`
}

// Message 消息模型
type Message struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ConversationID uuid.UUID `json:"conversation_id" gorm:"type:uuid;not null"`
	Role           string    `json:"role" gorm:"not null"` // "user" | "assistant"
	Content        string    `json:"content" gorm:"type:text;not null"`
	CreatedAt      time.Time `json:"created_at"`

	// 关联
	Conversation Conversation `json:"conversation,omitempty" gorm:"foreignKey:ConversationID"`
}

// File 文件模型
type File struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name      string    `json:"name" gorm:"not null"`
	Path      string    `json:"path" gorm:"not null"`
	Type      string    `json:"type"`
	Size      int64     `json:"size"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	CreatedAt time.Time `json:"created_at"`

	// 关联
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}

// Document RAG文档模型
type Document struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	Embedding []float64 `json:"-" gorm:"type:vector(1536)"` // OpenAI embedding 维度
	Source    string    `json:"source"`                     // 来源文件名
	Metadata  string    `json:"metadata" gorm:"type:jsonb"` // 额外元数据
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate 创建前设置UUID
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

func (c *Conversation) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}

func (f *File) BeforeCreate(tx *gorm.DB) error {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}
	return nil
}

func (d *Document) BeforeCreate(tx *gorm.DB) error {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}
	return nil
}

func SaveConversation() {}

func GetConversation() {}

func GetConversations() {}

func CreateConversation() {}
