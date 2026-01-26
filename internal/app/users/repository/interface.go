package repository

import "videocall/internal/app/users/schema"

type UsersRepositoryInterface interface {
	CreateUsers(users *schema.UserRequest) error
	UpdateUsers(id int, users *schema.UserRequest) error
	DeleteUsers(id int) error
	GetUserById(id int) (*schema.Users, error)
	GetAllUser() ([]*schema.Users, error)
}
