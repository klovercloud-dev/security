package service

import v1 "github.com/klovercloud-ci-cd/security/core/v1"

// Role business operations.
type Role interface {
	Store(role v1.RoleDto) error
	Get() []v1.RoleDto
	GetByName(name string) v1.RoleDto
	Delete(name string) error
	Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error
}
