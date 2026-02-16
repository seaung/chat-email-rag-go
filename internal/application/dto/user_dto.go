package dto

import "github.com/google/uuid"

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string      `json:"token"`
	User  UserInfoDTO `json:"user"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required,min=6"`
	Email    string      `json:"email" binding:"required,email"`
	Nickname string      `json:"nickname"`
	RoleIDs  []uuid.UUID `json:"role_ids"` // 管理员分配的角色
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Email    string      `json:"email" binding:"omitempty,email"`
	Nickname string      `json:"nickname"`
	Avatar   string      `json:"avatar"`
	Status   *int        `json:"status"` // 指针以区分0值和未传值
	RoleIDs  []uuid.UUID `json:"role_ids"`
}

// UserInfoDTO 用户信息响应（不含密码）
type UserInfoDTO struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Nickname string    `json:"nickname"`
	Avatar   string    `json:"avatar"`
	Status   int       `json:"status"`
	Roles    []RoleDTO `json:"roles"`
}

type RoleDTO struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"`
}
