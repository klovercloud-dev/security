package repository

import v1 "github.com/klovercloud-ci/core/v1"

// Permission Repository operations permission.
type Permission interface {
	Store(permission v1.Permission) error
	Get() []v1.Permission
	Delete(permissionName string) error
}
