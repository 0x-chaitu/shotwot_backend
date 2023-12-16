package service

import (
	"context"
	"errors"
	"fmt"
	"shotwot_backend/internal/domain"
	"shotwot_backend/internal/repository"
	"shotwot_backend/pkg/auth"
	"shotwot_backend/pkg/hash"
	"time"
)

type AccountsService struct {
	repo            repository.Accounts
	tokenManager    auth.TokenManager
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewAccountsService(repo repository.Accounts, tokenManager auth.TokenManager, accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration) *AccountsService {
	return &AccountsService{
		repo:            repo,
		tokenManager:    tokenManager,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (s *AccountsService) SignUp(ctx context.Context, input AccountSignUpInput) error {

	passwordHash, err := hash.Hash(input.Password)
	if err != nil {
		return err
	}

	account := &domain.Account{
		Email:    input.Email,
		Name:     input.Name,
		Password: passwordHash,
	}

	if err := s.repo.Create(ctx, account); err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return err
		}
		return err
	}
	return nil
}

func (s *AccountsService) SignIn(ctx context.Context, input AccountSignInInput) (*Tokens, error) {
	user, err := s.repo.GetByCredentials(ctx, input.Name, input.Password)
	if err != nil {
		return nil, err
	}
	token, err := s.tokenManager.NewJWT(fmt.Sprint(user.Id), s.accessTokenTTL)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		AccessToken: token,
	}, nil
}
