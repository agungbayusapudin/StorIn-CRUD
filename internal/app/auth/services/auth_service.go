package services

import (
	"videocall/internal/app/auth/repository"
	"videocall/internal/app/auth/schema"
	userSchema "videocall/internal/app/users/schema"
)

type authService struct {
	repo repository.AuthRepositoryInterface
}

func NewAuthService(repo repository.AuthRepositoryInterface) AuthServiceInterface {
	return &authService{repo: repo}
}

func (svc *authService) Login(userLoginReq *schema.UserLoginRequest) (*userSchema.Users, error) {
	return &userSchema.Users{}, nil
}

func (svc *authService) Register(userRegisterReq *schema.UserRegisterRequest) (*userSchema.Users, error) {
	return &userSchema.Users{}, nil
}

func (svc *authService) Logout(userID int) error {
	return nil
}

func (svc *authService) GenerateToken(userId int) (string, error) {
	return "", nil
}
