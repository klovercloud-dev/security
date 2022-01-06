package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type permissionService struct {
	repo repository.Permission
}

func (p permissionService) Store(permission v1.Permission) error {
	listOfPermission := p.Get()
	m := make(map[string]bool)

	for _, v := range listOfPermission {
		m[string(v.Name)] = true
	}

	if _, ok := m[string(permission.Name)]; !ok {
		return errors.New("Permission not valid!")
	}

	return p.repo.Store(permission)
}

func (p permissionService) Get() []v1.Permission {
	return p.repo.Get()
}

func (p permissionService) Delete(permissionName string) error {
	return p.repo.Delete(permissionName)
}

func NewPermissionService(repo repository.Permission) service.Permission {
	return &permissionService{
		repo: repo,
	}
}
