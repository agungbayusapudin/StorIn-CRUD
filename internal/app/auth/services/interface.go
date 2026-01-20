package services

import (
	"videocall/internal/app/auth/schema"
	userSchema "videocall/internal/app/users/schema"
)

type AuthServiceInterface interface {
	Login(userLoginRequest *schema.UserLoginRequest) (*userSchema.Users, error)
	Register(userRegisterRequest *schema.UserRegisterRequest) (*userSchema.Users, error)
	Logout(userID int) error
}
