package entity

import (
	"time"

	"github.com/google/uuid"
)

// User 用户聚合根
type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"` // 序列化时忽略密码
	Email        string    `json:"email"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Status       int       `json:"status"` // 1:正常, 0:禁用
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// 关联角色
	Roles []Role `json:"roles,omitempty"`
}

// Role 角色实体
type Role struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
