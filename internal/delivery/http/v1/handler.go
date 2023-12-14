package v1

import (
	"net/http"
	"shotwot_backend/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services
}

// Render for All Responses
func (rd *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Response is a wrapper response structure
type Response struct {
	Data interface{} `json:"data"`
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() http.Handler {
	r := chi.NewRouter()
	r.Mount("/users", h.initAccountsRoutes())

	return r
}
