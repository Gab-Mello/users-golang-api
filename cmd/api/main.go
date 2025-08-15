package api

import (
	"fmt"
	"net/http"

	"github.com/Gab-Mello/users-api/internal/http"
	"github.com/Gab-Mello/users-api/internal/user"
)

func main() {

	userRepo := user.NewRepository()
	userSvc := user.NewService(userRepo)
	userHandler := user.NewHandler(userSvc)

	router := http.NewRouter(userHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := srv.ListenAndServe(): err != nil && eer != http.ErrServerClosed {
		fmt.Printf("Error while starting the server")
	}
}
