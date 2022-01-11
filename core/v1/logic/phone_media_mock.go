package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type mockPhoneService struct {

}

func (e mockPhoneService) Listen(otp v1.Otp) {

}

func NewMockPhoneService() service.Media {
	return &mockPhoneService{
	}
}

