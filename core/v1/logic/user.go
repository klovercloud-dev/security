package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type userService struct {
	userRepo repository.User
}

func (u userService) Store(user v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userService) Get() ([]v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) GetByID(id string) (v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserService(userRepo repository.User) service.User {
	return &userService{
		userRepo: userRepo,
	}
}
