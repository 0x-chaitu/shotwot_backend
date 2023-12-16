package v1

import (
	"encoding/json"
	"errors"
	"net/http"
	"shotwot_backend/internal/domain"
	"shotwot_backend/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (h *Handler) initAccountsRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/signup", h.accountSignUp)
	r.Post("/signin", h.accountSignIn)
	return r

}

func (h *Handler) accountSignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var inp service.AccountSignUpInput
	err := decoder.Decode(&inp)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &Response{
			Data: "invalid input body",
		})
		return
	}
	if err := h.services.Accounts.SignUp(r.Context(), inp); err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			render.Status(r, http.StatusConflict)
			render.Render(w, r, &Response{
				Data: "User account already Exists",
			})
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, &Response{
			Data: err.Error(),
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &Response{
		Data: "signedup",
	})
}

func (h *Handler) accountSignIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var inp service.AccountSignInInput
	err := decoder.Decode(&inp)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.Render(w, r, &Response{
			Data: "invalid input body",
		})
		return
	}
	tokens, err := h.services.Accounts.SignIn(r.Context(), inp)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, &Response{
		Data: tokens,
	})
}
