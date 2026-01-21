package services

import (
	"videocall/internal/app/auth/schema"
	userSchema "videocall/internal/app/users/schema"
)

type AuthServiceInterface interface {
	Login(userLoginRequest *schema.UserLoginRequest) (*schema.UserAuthResponse, error)
	Register(userRegisterRequest *userSchema.UserRequest) error
	Logout(userID int) error
}
