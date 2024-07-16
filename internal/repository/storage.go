package repository

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"test/internal/config"
)

func InitDBConn(ctx context.Context, config *config.Config) *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PgUser, config.PgPassword, config.PgHost, config.PgPort, config.PgDatabaseName)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
