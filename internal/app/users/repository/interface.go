package repository

import "videocall/internal/app/users/schema"

type UsersRepositoryInterface interface {
	CreateUsers(users *schema.Users) error
	UpdateUsers(users *schema.UserRequest) error
	DeleteUsers(id int) error
	GetUserById(id int) (*schema.Users, error)
	GetAllUser() ([]*schema.Users, error)
}
