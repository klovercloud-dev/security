package repository

import (
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/enums"
)

// User Repository operations user.
type User interface {
	Store(user v1.User) error
	Get() []v1.User
	GetByID(id string) v1.User
	Delete(id string) error
	GetByEmail(email string) v1.User
	GetByPhone(phone string) v1.User
	GetByToken(token string) v1.User
	UpdatePassword(user v1.User) error
	UpdateUserResourcePermissionDto(id string, userResourcePermissionDto v1.UserResourcePermission) error
	AttachCompany(id, companyId string) error
	UpdateStatus(id string, status enums.STATUS) error
	GetUsersByCompanyId(companyId string, status enums.STATUS) []v1.User
}
