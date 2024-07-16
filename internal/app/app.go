package app

import (
	"context"
	"fmt"
	"net/http"
	"test/internal/config"
	delivery_http "test/internal/delivery/http"
	"test/internal/repository"
	"test/internal/usecase/user"
)

func Run() {
	config := config.Setup()

	ctx := context.Background()
	db := repository.InitDBConn(ctx, config)

	userRepository := repository.NewUserRepository(db)
	userUsecase := user.NewUserUsecase(userRepository)
	userHandler := delivery_http.NewUserHandler(userUsecase)

	http.HandleFunc("/api/register", userHandler.RegisterUser)

	http.ListenAndServe(fmt.Sprintf(":%s", config.BackendPort), nil)
}
