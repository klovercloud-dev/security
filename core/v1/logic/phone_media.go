package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type phoneService struct {

}

func (e phoneService) Listen(otp v1.Otp) {

}

func NewPhoneService() service.Media {
	return &phoneService{
	}
}

