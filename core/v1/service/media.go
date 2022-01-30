package service

import v1 "github.com/klovercloud-ci-cd/security/core/v1"

// Media business operations.
type Media interface {
	Listen(otp v1.Otp)
}
