package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type mockOtpService struct {
	repo repository.Otp
}

func (o mockOtpService) Store(otp v1.Otp) error {
	return o.repo.Store(otp)
}

func (o mockOtpService) FindByOtp(otp string) v1.Otp {
	return o.repo.FindByOtp(otp)
}

func (o mockOtpService) IsValid(otp string) bool {
	return true
}

func NewMockOtpService() service.Otp {
	return &mockOtpService{
	}
}
