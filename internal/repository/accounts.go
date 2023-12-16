package repository

import (
	"context"
	"shotwot_backend/internal/domain"
	postgres "shotwot_backend/pkg/database"
	"shotwot_backend/pkg/logger"
)

type AccountsRepo struct {
	*postgres.Postgres
}

func NewAccountsRepo(pg *postgres.Postgres) *AccountsRepo {
	return &AccountsRepo{pg}
}

func (r *AccountsRepo) Create(ctx context.Context, user *domain.Account) error {
	sql, args, err := r.Builder.
		Insert("user_account").
		Columns("username, password_hash, email").
		Values(user.Name, user.Password, user.Email).
		ToSql()
	if err != nil {
		logger.Errorf("Store - r.Builder: %v", err)
		return err
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		if postgres.UniqueKey(err) {
			return domain.ErrAccountAlreadyExists
		}

		logger.Errorf("Store - r.Pool.Exec: %v", err)
		return err
	}

	return nil
}

func (r *AccountsRepo) GetByCredentials(ctx context.Context, userIdentifier, password string) (*domain.Account, error) {
	sql, args, err := r.Builder.Select("id").From("user_account").Where("username = ?", userIdentifier).ToSql()
	if err != nil {
		logger.Errorf("Get - r.Builder: %v", err)
		return nil, err
	}
	var account domain.Account

	if err = r.Pool.QueryRow(ctx, sql, args...).Scan(&account.Id); err != nil {
		logger.Errorf(" - r.Pool.Exec: %v", err)
		return nil, err
	}
	return &account, nil
}
