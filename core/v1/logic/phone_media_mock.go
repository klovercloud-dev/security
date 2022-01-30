package logic

import (
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
)

type mockPhoneService struct {
}

func (e mockPhoneService) Listen(otp v1.Otp) {

}

// NewMockPhoneService returns service.Media type service
func NewMockPhoneService() service.Media {
	return &mockPhoneService{}
}
