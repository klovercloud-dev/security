package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Otp interface {
	// check if any record exists by ID
	// If any record exists, update that record by newly passed object
	// else store
	Store(otp v1.Otp) error
	FindByOtp(otp string) v1.Otp
}