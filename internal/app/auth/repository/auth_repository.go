package repository

import (
	"database/sql"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepositoryInterface {
	return &authRepository{db: db}
}

func (repo *authRepository) SaveUserSession(userId int, token string) error {
	return nil
}

func (repo *authRepository) GetUserSession(userId int) (string, error) {
	return "", nil
}

func (repo *authRepository) DeleteUserSession(userId int) error {
	return nil
}
