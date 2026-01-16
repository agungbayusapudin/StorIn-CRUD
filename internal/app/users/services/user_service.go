package services

import (
	"videocall/internal/app/users/repository"
	"videocall/internal/app/users/schema"
)

type userServices struct {
	repo repository.UsersRepositoryInterface
}

func NewUserServices(repo repository.UsersRepositoryInterface) UserServiceInterface {
	return &userServices{repo: repo}
}

func (serv *userServices) CreateUsers(users *schema.Users) error {
	return nil
}

func (serv *userServices) UpdateUsers(users *schema.UserRequest) error {
	return nil
}

func (serv *userServices) DeleteUsers(id int) error {
	return nil
}

func (serv *userServices) GetUserById(id int) (*schema.Users, error) {
	return nil, nil
}

func (serv *userServices) GetAllUser() ([]*schema.Users, error) {
	return nil, nil
}
