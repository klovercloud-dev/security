package repository

import v1 "github.com/klovercloud-ci/core/v1"

type UserResourcePermission interface {
	Store(userResourcePermission v1.UserResourcePermission) error
	Get() ([]v1.UserResourcePermission, error)
	GetByUserID(userID string) (v1.UserResourcePermission, error)
	Delete(userID string) error
	Update(userResourcePermission v1.UserResourcePermission) error
}
