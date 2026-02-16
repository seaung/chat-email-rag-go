package persistence

import (
	"chat-email-rag-go/internal/domain/entity"
	"chat-email-rag-go/internal/domain/repo"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	Nickname     string
	Avatar       string
	Status       int
	Roles        []RoleModel `gorm:"many2many:user_roles;"`
}

type RoleModel struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Code        string    `gorm:"uniqueIndex"`
	Name        string
	Description string
}

// 自动建表时映射
func (UserModel) TableName() string { return "users" }
func (RoleModel) TableName() string { return "roles" }

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) repo.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *entity.User) error {
	model := r.toModel(user)
	// GORM 会自动处理 many2many 关联
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var model UserModel
	err := r.db.WithContext(ctx).Preload("Roles").First(&model, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return r.toDomain(&model), nil
}

func (r *PostgresUserRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var model UserModel
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&model).Error
	if err != nil {
		return nil, err
	}
	return r.toDomain(&model), nil
}

func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var models []UserModel
	err := r.db.WithContext(ctx).Preload("Roles").Find(&models).Error
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, len(models))
	for i, m := range models {
		users[i] = *r.toDomain(&m)
	}
	return users, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *entity.User) error {
	model := r.toModel(user)
	// 更新基础信息
	return r.db.WithContext(ctx).Model(&model).Updates(map[string]interface{}{
		"email":    model.Email,
		"nickname": model.Nickname,
		"avatar":   model.Avatar,
		"status":   model.Status,
	}).Error
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, "id = ?", id).Error
}

func (r *PostgresUserRepository) FindRolesByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.Role, error) {
	var models []RoleModel
	err := r.db.WithContext(ctx).Find(&models, "id IN ?", ids).Error
	if err != nil {
		return nil, err
	}
	roles := make([]entity.Role, len(models))
	for i, m := range models {
		roles[i] = entity.Role{ID: m.ID, Code: m.Code, Name: m.Name, Description: m.Description}
	}
	return roles, nil
}

// 辅助转换方法
func (r *PostgresUserRepository) toModel(u *entity.User) *UserModel {
	return &UserModel{
		ID:           u.ID,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		Email:        u.Email,
		Nickname:     u.Nickname,
		Avatar:       u.Avatar,
		Status:       u.Status,
	}
}

func (r *PostgresUserRepository) toDomain(m *UserModel) *entity.User {
	roles := make([]entity.Role, len(m.Roles))
	for i, rm := range m.Roles {
		roles[i] = entity.Role{ID: rm.ID, Code: rm.Code, Name: rm.Name}
	}
	return &entity.User{
		ID:           m.ID,
		Username:     m.Username,
		PasswordHash: m.PasswordHash,
		Email:        m.Email,
		Nickname:     m.Nickname,
		Avatar:       m.Avatar,
		Status:       m.Status,
		Roles:        roles,
	}
}
