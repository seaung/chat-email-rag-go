package middlewares

import (
	"chat-email-rag-go/internal/domain/entity"
	"chat-email-rag-go/internal/domain/repo"
	"chat-email-rag-go/internal/infrastructure/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthMiddleware 认证中间件
type AuthMiddleware struct {
	jwtUtil  *auth.JWTUtil
	userRepo repo.UserRepository
}

// NewAuthMiddleware 构造函数
func NewAuthMiddleware(jwtUtil *auth.JWTUtil, userRepo repo.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil:  jwtUtil,
		userRepo: userRepo,
	}
}

// RequireAuth 验证身份的中间件函数
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取 Authorization Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// 2. 解析 Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization format is invalid (Bearer <token>)"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. 验证 Token
		claims, err := m.jwtUtil.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 4. 从数据库查询用户最新信息 (确保用户未被删除或封禁)
		userUUID, err := uuid.Parse(claims.UserID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
			c.Abort()
			return
		}

		user, err := m.userRepo.FindByID(c.Request.Context(), userUUID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// 5. 检查用户状态
		if user.Status == 0 {
			c.JSON(http.StatusForbidden, gin.H{"error": "User account is banned"})
			c.Abort()
			return
		}

		// 6. 将用户实体存入 Context，供后续 Handler 使用
		c.Set("currentUser", user)

		c.Next()
	}
}

// GetCurrentUser 从 Context 获取当前登录用户的辅助函数
func GetCurrentUser(ctx *gin.Context) (*entity.User, bool) {
	if user, exists := ctx.Get("currentUser"); exists {
		if u, ok := user.(*entity.User); ok {
			return u, true
		}
	}
	return nil, false
}
