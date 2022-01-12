package logic

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
	"io"
	"net/mail"
)

type userService struct {
	userRepo repository.User
	urpService  service.UserResourcePermission
	tokenService service.Token
	otpService service.Otp
	emailMediaService service.Media
	phoneMediaService service.Media
	httpClientService service.HttpClient
}

func (u userService) AttachCompany(company v1.Company, companyId,token string) error {
	tokenObject:=u.tokenService.GetByToken(token)
	if tokenObject.Uid==""{
		return errors.New("no token found")
	}
	marshal, marshalErr := json.Marshal(company)
	if marshalErr != nil {
		return marshalErr
	}
	header := make(map[string]string)
	header["token"] = token
	header["Authorization"]="Bearer "+token
	header["Content-Type"] = "application/json"
	_, err := u.httpClientService.Post(config.ApiServerUrl+"/companies", header, marshal)
	if err != nil {
		return  err
	}
	return u.userRepo.AttachCompany(tokenObject.Uid,companyId)
}

func (u userService) GetByOtp(otp string) v1.User {
	otpObject:=u.otpService.FindByOtp(otp)
	return u.GetByID(otpObject.ID)
}

func (u userService) GetByPhone(phone string) v1.User {
	return u.userRepo.GetByPhone(phone)
}

func (u userService) SendOtp(email, phone string) error {
	var user v1.User
	if email!=""{
		user=u.GetByEmail(email)
	}else if phone!=""{
		user=u.GetByPhone(phone)
	}
	otp:=v1.Otp{
		ID:    user.ID,
		Email: user.Email,
		Phone: user.Phone,
		Otp:   u.generateOtp(6),
	}
	if email!=""{
		u.emailMediaService.Listen(otp)
	}else{
		 u.phoneMediaService.Listen(otp)
	}
	return u.otpService.Store(otp)
}

func (u userService) generateOtp(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func (u userService) UpdatePassword(user v1.User) error {
	return u.userRepo.UpdatePassword(user)
}



func (u userService) UpdateToken(token, refreshToken, existingToken string) error {
	return u.tokenService.Update(token,refreshToken,existingToken)
}

func (u userService) GetByEmail(email string) v1.User {
	return u.userRepo.GetByEmail(email)
}

func (u userService) Store(userWithResourcePermission v1.UserRegistrationDto) error {
	user, userResourcePermission := v1.GetUserAndResourcePermissionBody(userWithResourcePermission)
	isUserExist := u.userRepo.GetByEmail(user.Email)
	if isUserExist.Email != "" {
		return errors.New("email is already registered")
	}

	err:=userWithResourcePermission.Validate()
	if err != nil {
		return err
	}

	err = u.userRepo.Store(user)
	if err != nil {
		return err
	}

	err = u.urpService.Store(userResourcePermission)
	if err != nil {
		return err
	}
	return nil
}

func (u userService) Get() []v1.User {
	users := u.userRepo.Get()
	return users
}

func (u userService) GetByID(id string) v1.User {
	return u.userRepo.GetByID(id)
}

func (u userService) Delete(id string) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func mailValidation(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewUserService(userRepo repository.User, urpService service.UserResourcePermission,tokenService service.Token,otpService service.Otp,emailMediaService service.Media,phoneMediaService service.Media, httpClientService service.HttpClient) service.User {
	return &userService{
		userRepo:          userRepo,
		urpService:        urpService,
		tokenService:      tokenService,
		otpService:        otpService,
		emailMediaService: emailMediaService,
		phoneMediaService: phoneMediaService,
		httpClientService: httpClientService,
	}
}
