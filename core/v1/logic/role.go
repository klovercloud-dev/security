package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
)

type roleService struct {
	roleRepo       repository.Role
	permissionRepo repository.Permission
}

func (r roleService) Get() ([]v1.Role, int64) {
	roles, total := r.roleRepo.Get()
	return roles, total
}

func (r roleService) Store(role v1.Role) error {
	roles, _ := r.roleRepo.Get()
	for _, each := range roles {
		if each.Name == role.Name {
			return errors.New("Role already exists!")
		}
	}
	err := r.roleRepo.Store(role)
	if err != nil {
		return err
	}
	return nil
}

func (r roleService) GetByName(name string) (v1.Role, error) {
	roleByName, err := r.roleRepo.GetByName(name)
	if err != nil {
		return v1.Role{}, err
	}
	return roleByName, nil
}

func (r roleService) Delete(name string) error {
	err := r.roleRepo.Delete(name)
	if err != nil {
		return err
	}
	return nil
}

func (r roleService) Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error {
	roles, _ := r.roleRepo.Get()
	for _, each := range roles {
		if each.Name != name {
			return errors.New("Role already exists!")
		}
	}

	m := make(map[string]bool)

	listOfPermissions := r.permissionRepo.Get()
	for _, each := range listOfPermissions {
		m[string(each.Name)] = true
	}

	for _, each := range permissions {
		if _, ok := m[string(each.Name)]; !ok {
			return errors.New("Permission not found!")
		}
	}
	if option.Option == enums.APPEND_PERMISSION {
		err := r.roleRepo.AppendPermissions(name, permissions)
		if err != nil {
			return err
		}
	} else if option.Option == enums.REMOVE_PERMISSION {
		err := r.roleRepo.RemovePermissions(name, permissions)
		if err != nil {
			return err
		}
	} else {
		err := r.roleRepo.Update(name, permissions)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewRoleService(roleRepo repository.Role) service.Role {
	return &roleService{roleRepo: roleRepo}
}
