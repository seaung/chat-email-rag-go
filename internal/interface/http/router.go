package http

import (
	"chat-email-rag-go/internal/application/service"
	"chat-email-rag-go/internal/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, userService *service.UserAppService) {
	userHandler := handler.NewUserHandler(userService)

	// API Group
	api := r.Group("/api/v1")
	{
		api.POST("/login", userHandler.Login)

		adminGroup := api.Group("/admin/users")
		// adminGroup.Use(middleware.AuthMiddleware()) // 实际项目中请解开注释
		{
			adminGroup.POST("", userHandler.CreateUser)
			adminGroup.GET("", userHandler.ListUsers)
			adminGroup.GET("/:id", userHandler.GetUser)
			adminGroup.PUT("/:id", userHandler.UpdateUser)
			adminGroup.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}
