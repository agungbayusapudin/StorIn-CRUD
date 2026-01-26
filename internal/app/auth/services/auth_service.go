package services

import (
	"videocall/internal/app/auth/schema"
	"videocall/internal/app/users/repository"
	userSchema "videocall/internal/app/users/schema"
	"videocall/pkg/jwt"
)

type authService struct {
	user repository.UsersRepositoryInterface
}

func NewAuthService(user repository.UsersRepositoryInterface, secret string) AuthServiceInterface {
	return &authService{
		user: user,
	}
}

func (svc *authService) Login(userLoginReq *schema.UserLoginRequest) (*schema.UserAuthResponse, error) {
	userData, err := svc.user.GetUserById(userLoginReq.Id)
	if err != nil {
		return nil, err
	}

	TokenDataRequest := schema.TokenRequest{
		Id:       userData.ID,
		Username: userData.Username,
		Role:     "User",
	}

	token, err := jwt.CreateToken(&TokenDataRequest)

	responseData := schema.UserAuthResponse{
		Token:    token,
		UserData: userData,
	}

	return &responseData, nil
}

func (svc *authService) Register(userRegisterReq *userSchema.UserRequest) error {
	err := svc.user.CreateUsers(userRegisterReq)
	if err != nil {
		return err
	}

	return nil

}

func (svc *authService) Logout(userID int) error {
	// logout haurs menggunakan storage sementara
	return nil
}
