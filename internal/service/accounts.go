package service

import (
	"context"
	"shotwot_backend/internal/repository"
)

type AccountsService struct {
	repo *repository.Accounts
}

func NewAccountsService(repo repository.Accounts) *AccountsService {
	return &AccountsService{
		repo: &repo,
	}
}

func (s *AccountsService) SignUp(ctx context.Context, input AccountSignUpInput) error {
	return nil
}

func (s *AccountsService) SignIn(ctx context.Context, input AccountSignInInput) (Tokens, error) {
	return Tokens{}, nil
}
