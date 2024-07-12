package app

import (
	"context"
	"net/http"
	delivery_http "test/internal/delivery/http"
	"test/internal/repository"
	"test/internal/usecase/user"
)

func Run() {
	ctx := context.Background()
	db := repository.InitDBConn(ctx)

	userRepository := repository.NewUserRepository(db)
	userUsecase := user.NewUserUsecase(userRepository)
	userHandler := delivery_http.NewUserHandler(userUsecase)

	http.HandleFunc("/api/register", userHandler.RegisterUser)

	http.ListenAndServe(":8080", nil)
}
