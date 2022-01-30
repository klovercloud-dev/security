package service

import v1 "github.com/klovercloud-ci-cd/security/core/v1"

// UserResourcePermission business operations.
type UserResourcePermission interface {
	GetByUserID(userID string) v1.UserResourcePermissionDto
}
