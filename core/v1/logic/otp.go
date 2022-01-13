package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"time"
)

type otpService struct {
	repo repository.Otp
}

func (o otpService) Store(otp v1.Otp) error {
	return o.repo.Store(otp)
}

func (o otpService) FindByOtp(otp string) v1.Otp {
	return o.repo.FindByOtp(otp)
}

func (o otpService) IsValid(otpStr string) bool {
	otp := o.repo.FindByOtp(otpStr)
	if otp.ID == "" {
		return false
	}
	if time.Now().UTC().After(otp.Exp) {
		return false
	}
	return true
}

func NewOtpService(repo repository.Otp) service.Otp {
	return &otpService{
		repo: repo,
	}
}
