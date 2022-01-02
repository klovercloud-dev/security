package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Permission interface {
	Store(permission v1.Permission) error
	Get() ([]v1.Permission, error)
	Delete(permissionName string) error
}
