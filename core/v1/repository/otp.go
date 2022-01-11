package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Otp interface {
	Store(otp v1.Otp) error
	FindByOtp(otp string ) v1.Otp
}

