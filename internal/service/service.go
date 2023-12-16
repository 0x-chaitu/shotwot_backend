package service

import (
	"context"
	"shotwot_backend/internal/repository"
	"shotwot_backend/pkg/auth"
	"time"
)

type AccountSignUpInput struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AccountSignInInput struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Tokens struct {
	AccessToken  string `json:"accesstoken"`
	RefreshToken string `json:"refreashtoken"`
}

type Accounts interface {
	SignUp(ctx context.Context, input AccountSignUpInput) error
	SignIn(ctx context.Context, input AccountSignInInput) (*Tokens, error)
	// RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
	// Verify(ctx context.Context, userID primitive.ObjectID, hash string) error
}

type Services struct {
	Accounts Accounts
}

type Deps struct {
	Repos        *repository.Repositories
	TokenManager auth.TokenManager

	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

func NewServices(deps Deps) *Services {
	accountService := NewAccountsService(deps.Repos.Accounts, deps.TokenManager, deps.AccessTokenTTL, deps.RefreshTokenTTL)
	return &Services{
		Accounts: accountService,
	}
}
