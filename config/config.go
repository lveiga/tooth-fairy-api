package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// AppConfig - responsible to manage application configuration
type AppConfig struct {
	BindAddr             string
	DatabaseHost         string
	DatabasePort         string
	DatabaseName         string
	DatabaseUser         string
	DatabasePass         string
	Environment          string
	ContextGenricTimeout time.Duration
}

// New - responsible to store env configs
func New() *AppConfig {
	if err := godotenv.Load(".env"); err != nil {
		//fmt.Println("WARN - ERROR TO LOAD .ENV FILE")
	}

	return &AppConfig{
		BindAddr:             os.Getenv("BIND_ADDR"),
		DatabaseHost:         os.Getenv("DATABASE_HOST"),
		DatabasePort:         os.Getenv("DATABASE_PORT"),
		DatabaseName:         os.Getenv("DATABASE_NAME"),
		DatabaseUser:         os.Getenv("DATABASE_USER"),
		DatabasePass:         os.Getenv("DATABASE_PASS"),
		Environment:          os.Getenv("ENVIRONMENT"),
		ContextGenricTimeout: 5 * time.Second,
	}
}
