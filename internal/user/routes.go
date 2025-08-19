package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.Create)
	r.Post("/", h.Get)
	r.Get("/{id}", h.List)
	r.Delete("/{id}", h.Delete)

	return r
}
