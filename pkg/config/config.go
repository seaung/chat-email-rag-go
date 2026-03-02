package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBURI          string
	OllamaURI      string
	OllmaBaseModel string
	OllamaEmb      string
	JWTSecRet      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
	}

	return &Config{
		Port:           getEnv("PORT", "9527"),
		DBURI:          getEnv("DBURI", ""),
		OllamaURI:      getEnv("OLLAMA_URI", ""),
		OllmaBaseModel: getEnv("OLLAMA_MODEL", ""),
		OllamaEmb:      getEnv("OLLAMA_EMB", ""),
		JWTSecRet:      getEnv("JWT_SEC_RET", ""),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
