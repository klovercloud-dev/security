package service

import v1 "github.com/klovercloud-ci/core/v1"

// Permission business operations.
type Permission interface {
	Store(permission v1.Permission) error
	Get() []v1.Permission
	Delete(permissionName string) error
}
