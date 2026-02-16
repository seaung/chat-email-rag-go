package service

import (
	"chat-email-rag-go/internal/application/dto"
	"chat-email-rag-go/internal/domain/entity"
	"chat-email-rag-go/internal/domain/repo"
	"chat-email-rag-go/internal/infrastructure/auth"
	"context"
	"errors"

	"github.com/google/uuid"
)

type UserAppService struct {
	userRepo repo.UserRepository
	jwtUtil  *auth.JWTUtil
}

func NewUserAppService(userRepo repo.UserRepository, jwtUtil *auth.JWTUtil) *UserAppService {
	return &UserAppService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

// Login 用户登录
func (s *UserAppService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !auth.CheckPassword(req.Password, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}

	if user.Status == 0 {
		return nil, errors.New("user is banned")
	}

	token, err := s.jwtUtil.GenerateToken(user.ID.String(), user.Username)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User:  *s.toDTO(user),
	}, nil
}

// CreateUser 创建用户（管理员功能）
func (s *UserAppService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) error {
	// 检查用户名是否存在
	_, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err == nil {
		return errors.New("username already exists")
	}

	hashedPwd, err := auth.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &entity.User{
		ID:           uuid.New(),
		Username:     req.Username,
		PasswordHash: hashedPwd,
		Email:        req.Email,
		Nickname:     req.Nickname,
		Status:       1, // 默认启用
	}

	// 处理角色分配
	if len(req.RoleIDs) > 0 {
		roles, err := s.userRepo.FindRolesByIDs(ctx, req.RoleIDs)
		if err != nil {
			return err
		}
		user.Roles = roles
	}

	return s.userRepo.Create(ctx, user)
}

// UpdateUser 更新用户
func (s *UserAppService) UpdateUser(ctx context.Context, id uuid.UUID, req *dto.UpdateUserRequest) error {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}

	// 更新字段
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Status != nil {
		user.Status = *req.Status
	}

	// 处理角色更新
	if req.RoleIDs != nil {
		roles, err := s.userRepo.FindRolesByIDs(ctx, req.RoleIDs)
		if err != nil {
			return err
		}
		user.Roles = roles
	}

	return s.userRepo.Update(ctx, user)
}

// DeleteUser 删除用户
func (s *UserAppService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.userRepo.Delete(ctx, id)
}

// GetUser 获取用户详情
func (s *UserAppService) GetUser(ctx context.Context, id uuid.UUID) (*dto.UserInfoDTO, error) {
	user, err := s.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return s.toDTO(user), nil
}

// ListUsers 获取用户列表
func (s *UserAppService) ListUsers(ctx context.Context) ([]dto.UserInfoDTO, error) {
	users, err := s.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	dtos := make([]dto.UserInfoDTO, len(users))
	for i, u := range users {
		dtos[i] = *s.toDTO(&u)
	}
	return dtos, nil
}

// 辅助转换
func (s *UserAppService) toDTO(user *entity.User) *dto.UserInfoDTO {
	roles := make([]dto.RoleDTO, len(user.Roles))
	for i, r := range user.Roles {
		roles[i] = dto.RoleDTO{ID: r.ID, Code: r.Code, Name: r.Name}
	}
	return &dto.UserInfoDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Status:   user.Status,
		Roles:    roles,
	}
}
