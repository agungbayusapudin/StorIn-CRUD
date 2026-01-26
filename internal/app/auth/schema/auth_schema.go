package schema

import "videocall/internal/app/users/schema"

type UserLoginRequest struct {
	Id       int     `json:"id"`
	Username *string `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email"`
}

type UserRegisterRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
	Phone     *string `json:"phone"`
}

type UserAuthResponse struct {
	Token    string        `json:"token"`
	UserData *schema.Users `json:"user"`
}

type UserSession struct {
	ID int `json:"id"`
}

type TokenRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
