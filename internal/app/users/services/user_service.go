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
	err := serv.repo.CreateUsers(users)
	if err != nil {
		return err
	}
	return nil
}

func (serv *userServices) UpdateUsers(id int, users *schema.UserRequest) error {
	err := serv.repo.UpdateUsers(id, users)
	if err != nil {
		return err
	}
	return nil
}

func (serv *userServices) DeleteUsers(id int) error {
	err := serv.repo.DeleteUsers(id)
	if err != nil {
		return err
	}
	return nil
}

func (serv *userServices) GetUserById(id int) (*schema.Users, error) {
	user, err := serv.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (serv *userServices) GetAllUser() ([]*schema.Users, error) {
	user, err := serv.repo.GetAllUser()
	if err != nil {
		return nil, err
	}

	return user, nil
}
