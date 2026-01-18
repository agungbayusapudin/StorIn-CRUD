package services

import "videocall/internal/app/users/schema"

type UserServiceInterface interface {
	CreateUsers(users *schema.Users) error
	UpdateUsers(id int, users *schema.UserRequest) error
	DeleteUsers(id int) error
	GetUserById(id int) (*schema.Users, error)
	GetAllUser() ([]*schema.Users, error)
}
