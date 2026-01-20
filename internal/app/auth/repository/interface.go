package repository

type AuthRepositoryInterface interface {
	SaveUserSession(userID int, token string) error
	GetUserSession(userID int) (string, error)
	DeleteUserSession(userID int) error
}
