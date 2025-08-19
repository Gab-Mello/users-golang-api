package user

import "net/http"

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {

}
