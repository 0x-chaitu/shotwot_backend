package repository

import (
	"context"
	"shotwot_backend/internal/domain"
	postgres "shotwot_backend/pkg/database"
)

type Accounts interface {
	Create(ctx context.Context, user domain.Account) error

	//userIdentifier can be username, email
	GetByCredentials(ctx context.Context, userIdentifier, password string) (domain.Account, error)
}

type Repositories struct {
	Accounts Accounts
}

func NewRepositories(db *postgres.Postgres) *Repositories {
	return &Repositories{
		Accounts: NewAccountsRepo(db),
	}
}
