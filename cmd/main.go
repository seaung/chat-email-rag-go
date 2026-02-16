package main

import (
	"chat-email-rag-go/internal/application/service"
	"chat-email-rag-go/internal/infrastructure/auth"
	"chat-email-rag-go/internal/infrastructure/persistence"
	"chat-email-rag-go/internal/interface/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123456 dbname=rag_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	urepo := persistence.NewPostgresUserRepository(db)
	jwtUtil := auth.NewJWTUtil("rag_application")
	uService := service.NewUserAppService(urepo, jwtUtil)

	r := gin.Default()
	http.SetupRouter(r, uService)
	r.Run(":9090")
}
