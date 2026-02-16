package repo

import (
	"chat-email-rag-go/internal/domain/entity"
	"context"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindAll(ctx context.Context) ([]entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	// 根据ID列表查询角色（用于关联）
	FindRolesByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.Role, error)
}
