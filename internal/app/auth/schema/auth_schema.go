package schema

type UserLoginRequest struct {
	ID int `json:"id"`
}

type UserRegisterRequest struct {
	ID int `json:"id"`
}

type UserSession struct {
	ID int `json:"id"`
}

type TokenRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
