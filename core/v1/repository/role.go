package repository

import v1 "github.com/klovercloud-ci/core/v1"

// Role Repository operations role.
type Role interface {
	Store(role v1.Role) error
	Get() []v1.Role
	GetByName(name string) v1.Role
	Delete(roleName string) error
	Update(name string, permissions []v1.Permission) error
	AppendPermissions(name string, permissions []v1.Permission) error
	RemovePermissions(name string, permissions []v1.Permission) error
}
