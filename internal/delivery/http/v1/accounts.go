package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) initAccountsRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/signup", h.accountSignUp)
	return r

}

type accountSignUpInput struct {
	Name     string `json:"username" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func (h *Handler) accountSignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var inp accountSignUpInput
	err := decoder.Decode(&inp)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &Response{
			Data: "invalid input body",
		})
		return
	}
	// if err := h.services.Accounts.SignUp(r.Context(),service.AccountSignUpInput{
	// 	Name: ,
	// })
	render.Status(r, http.StatusOK)
	render.Render(w, r, &Response{
		Data: "signedup",
	})
}
