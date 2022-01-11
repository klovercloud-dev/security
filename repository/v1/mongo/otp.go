package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"time"
)

// OtpCollection collection name
var (
	OtpCollection = "otpCollection"
)

type otpRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (o otpRepository) Store(otp v1.Otp) error {
	panic("implement me")
}

func (o otpRepository) FindByOtp(otp string) v1.Otp {
	panic("implement me")
}

func NewOtpRepository(timeout int) repository.Otp {
	return &otpRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
