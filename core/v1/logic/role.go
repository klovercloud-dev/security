package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type roleService struct {
	roleRepo repository.Role
}

func (r roleService) Get() ([]v1.Role, int64) {
	//TODO implement me
	panic("implement me")
}

func (r roleService) Store(role v1.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r roleService) GetByName(name string) (v1.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r roleService) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

func (r roleService) Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error {
	//TODO implement me
	panic("implement me")
}

func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{roleRepo: roleRepo}
}
