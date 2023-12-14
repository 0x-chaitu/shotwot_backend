package http

import (
	"net/http"
	"shotwot_backend/internal/config"
	v1 "shotwot_backend/internal/delivery/http/v1"
	"shotwot_backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	h.initAPI(r)

	return r
}

func (h *Handler) initAPI(r chi.Router) {
	handlerV1 := v1.NewHandler(h.services)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", handlerV1.Init())
	})
}
