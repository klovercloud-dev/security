package service

import v1 "github.com/klovercloud-ci/core/v1"

type Media interface {
	Listen(otp v1.Otp)
}
