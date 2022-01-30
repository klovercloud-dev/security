package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/repository"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"github.com/klovercloud-ci-cd/security/enums"
)

type roleService struct {
	roleRepo          repository.Role
	permissionService service.Permission
}

func (r roleService) Get() []v1.Role {
	roles := r.roleRepo.Get()
	return roles
}

func (r roleService) Store(role v1.Role) error {
	roles := r.roleRepo.Get()
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

func (r roleService) GetByName(name string) v1.Role {
	roleByName := r.roleRepo.GetByName(name)
	return roleByName
}

func (r roleService) Delete(name string) error {
	err := r.roleRepo.Delete(name)
	if err != nil {
		return err
	}
	return nil
}

func (r roleService) Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error {
	roles := r.roleRepo.Get()
	flag := false
	for _, each := range roles {
		if each.Name == name {
			flag = true
		}
	}
	if !flag {
		return errors.New("role does not exists")
	}

	m := make(map[string]bool)

	listOfPermissions := r.permissionService.Get()
	for _, each := range listOfPermissions {
		m[string(each.Name)] = true
	}

	for _, each := range permissions {
		if _, ok := m[string(each.Name)]; !ok {
			return errors.New("permission not found")
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

// NewRoleService returns service.Role type service
func NewRoleService(roleRepo repository.Role, permissionService service.Permission) service.Role {
	return &roleService{
		roleRepo:          roleRepo,
		permissionService: permissionService,
	}
}
