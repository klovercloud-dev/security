package logic

import (
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"log"
	"net/smtp"
)

type emailService struct {

}

func (e emailService) Listen(otp v1.Otp) {
	message:=`Hi `+otp.Email+`,`+`
		  Please find your OTP attached below. It will be expired within 5 minutes.
		  OTP:`+otp.Otp
	// Create authentication
	auth := smtp.PlainAuth("", config.MailServerHostEmail, config.MailServerHostEmailSecret, config.SmtpHost)
	// Send actual message
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.MailServerHostEmail, []string{otp.Email},[]byte (message))
	if err != nil {
		log.Println(err.Error())
	}
}

func NewEmailService() service.Media {
	return &emailService{
	}
}

