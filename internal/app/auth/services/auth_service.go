package services

import (
	"videocall/internal/app/auth/schema"
	"videocall/internal/app/users/repository"
	userSchema "videocall/internal/app/users/schema"
)

type authService struct {
	user repository.UsersRepositoryInterface

	// perlu token dan jwt secret
}

func NewAuthService(user repository.UsersRepositoryInterface) AuthServiceInterface {
	return &authService{user: user}
}

func (svc *authService) Login(userLoginReq *schema.UserLoginRequest) (*userSchema.Users, error) {
	return nil, nil
}

func (svc *authService) Register(userRegisterReq *schema.UserRegisterRequest) (*userSchema.Users, error) {
	return nil, nil
}

func (svc *authService) Logout(userID int) error {
	return nil
}
