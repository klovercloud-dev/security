package service

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/enums"
)

type User interface {
	Store(user v1.UserRegistrationDto) error
	Get() []v1.User
	GetByID(id string) v1.User
	Delete(id string) error
	GetByEmail(email string) v1.User
	GetByPhone(phone string) v1.User
	GetByOtp(otp string) v1.User
	UpdateToken(token, refreshToken, existingToken string) error
	UpdatePassword(user v1.User) error
	SendOtp(email, phone string) error
	AttachCompany(company v1.Company, companyId,token string) error
	UpdateStatus(id string,status enums.STATUS) error
	GetUsersByCompanyId(companyId string, status enums.STATUS) [] v1.User
}
