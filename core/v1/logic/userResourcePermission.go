package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type userResourcePermissionService struct {
	repo v1.UserResourcePermission
}

func (u userResourcePermissionService) Store(userResourcePermission v1.UserResourcePermission) error {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionService) Get() ([]v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionService) GetByUserID(userID string) (v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionService) Delete(userID string) error {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionService) Update(userResourcePermission v1.UserResourcePermission) error {
	//TODO implement me
	panic("implement me")
}

func NewUserResourcePermissionService(repo v1.UserResourcePermission) service.UserResourcePermission {
	return &userResourcePermissionService{
		repo: repo,
	}
}
