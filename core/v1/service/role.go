package service

import v1 "github.com/klovercloud-ci/core/v1"

type Role interface {
	Store(role v1.Role) error
	Get() []v1.Role
	GetByName(name string) v1.Role
	Delete(name string) error
	Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error
}
