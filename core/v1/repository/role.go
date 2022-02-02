package repository

import v1 "github.com/klovercloud-ci-cd/security/core/v1"

// Role Repository operations role.
type Role interface {
	Store(role v1.RoleDto) error
	Get() []v1.RoleDto
	GetByName(name string) v1.RoleDto
	Delete(roleName string) error
	Update(name string, permissions []v1.Permission) error
	AppendPermissions(name string, permissions []v1.Permission) error
	RemovePermissions(name string, permissions []v1.Permission) error
}
