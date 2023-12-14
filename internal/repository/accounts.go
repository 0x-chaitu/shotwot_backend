package repository

import (
	"context"
	"shotwot_backend/internal/domain"
	postgres "shotwot_backend/pkg/database"
)

type AccountsRepo struct {
	*postgres.Postgres
}

func NewAccountsRepo(pg *postgres.Postgres) *AccountsRepo {
	return &AccountsRepo{pg}
}

func (r *AccountsRepo) Create(ctx context.Context, user domain.Account) error {
	return nil
}

func (r *AccountsRepo) GetByCredentials(ctx context.Context, userIdentifier, password string) (domain.Account, error) {
	return domain.Account{}, nil
}
