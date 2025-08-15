package http

import (
	"net/http"

	"github.com/Gab-Mello/users-api/internal/user"
	"github.com/go-chi/chi/v5"
)

func NewRouter(uh *user.Handler) http.Handler {
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		user.RegisterRoutes(r, uh)
	})

	return r
}
