package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type userResourcePermissionService struct {
	repo repository.UserResourcePermission
}

func (u userResourcePermissionService) Store(userResourcePermission v1.UserResourcePermission) error {
	return u.repo.Store(userResourcePermission)
}

func (u userResourcePermissionService) Get() []v1.UserResourcePermission {
	return u.repo.Get()
}

func (u userResourcePermissionService) GetByUserID(userID string) v1.UserResourcePermission {
	return u.repo.GetByUserID(userID)
}

func (u userResourcePermissionService) Delete(userID string) error {
	return u.repo.Delete(userID)
}

func (u userResourcePermissionService) Update(userResourcePermission v1.UserResourcePermission) error {
	return u.repo.Update(userResourcePermission)
}

func NewUserResourcePermissionService(repo repository.UserResourcePermission) service.UserResourcePermission {
	return &userResourcePermissionService{
		repo: repo,
	}
}
