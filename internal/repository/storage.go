package repository

import (
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDBConn(ctx context.Context) *gorm.DB {
	dbURL := "postgres://user:password@pgsql:5432/app"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
