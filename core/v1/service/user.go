package service

import v1 "github.com/klovercloud-ci/core/v1"

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
	AttachCompany(id ,companyId string) error
}
