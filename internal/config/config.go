package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PgHost         string
	PgDatabaseName string
	PgUser         string
	PgPassword     string
	PgPort         string
	BackendPort    string
}

func env(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func Setup() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic("env file not found")
	}

	return &Config{
		PgHost:         env("PG_HOST", "pgsql"),
		PgDatabaseName: env("PG_DATABASE_NAME", "app"),
		PgUser:         env("PG_USER", "user"),
		PgPassword:     env("PG_PASSWORD", "password"),
		PgPort:         env("PG_PORT", "5432"),
		BackendPort:    env("BACKEND_PORT", "8080"),
	}
}
