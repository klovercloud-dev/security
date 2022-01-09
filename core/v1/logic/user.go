package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"net/mail"
)

type userService struct {
	userRepo repository.User
	urpRepo  repository.UserResourcePermission
}

func (u userService) UpdateToken(token, refreshToken, existingToken string) error {
	panic("implement me")
}

func (u userService) GetByEmail(email string) v1.User {
	return u.userRepo.GetByEmail(email)
}

func (u userService) Store(userWithResourcePermission v1.UserRegistrationDto) error {
	user, userResourcePermission := v1.GetUserAndResourcePermissionBody(userWithResourcePermission)
	mailFlag := mailValidation(userWithResourcePermission.Email)
	if mailFlag == false {
		return errors.New("email is not valid")
	}
	isUserExist := u.userRepo.GetByEmail(user.Email)
	if isUserExist.Email != "" {
		return errors.New("email is already registered")
	}
	if userWithResourcePermission.Password == "" {
		return errors.New("password is empty")
	} else if len(userWithResourcePermission.Password) < 8 {
		return errors.New("password is minimum 8 characters")
	}
	err := u.userRepo.Store(user)
	if err != nil {
		return err
	}

	err = u.urpRepo.Store(userResourcePermission)
	if err != nil {
		return err
	}
	return nil
}

func (u userService) Get() []v1.User {
	users := u.userRepo.Get()
	return users
}

func (u userService) GetByID(id string) (v1.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return v1.User{}, err
	}
	return user, nil
}

func (u userService) Delete(id string) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func mailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewUserService(userRepo repository.User, urpRepo repository.UserResourcePermission) service.User {
	return &userService{
		userRepo: userRepo,
		urpRepo:  urpRepo,
	}
}
