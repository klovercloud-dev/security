package logic

import (
	"crypto/rand"
	"errors"
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
}

func (u userService) AttachCompany(id, companyId string) error {
	panic("implement me")
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
	mailFlag := mailValidation(userWithResourcePermission.Email)
	if mailFlag == false {
		return errors.New("email is not valid")
	}
	isUserExist := u.userRepo.GetByEmail(user.Email)
	if isUserExist.Email != "" {
		return errors.New("email is already registered")
	}
	if userWithResourcePermission.Password == "" {
		return errors.New("password is empty")
	} else if len(userWithResourcePermission.Password) < 8 {
		return errors.New("password is minimum 8 characters")
	}
	err := u.userRepo.Store(user)
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

func NewUserService(userRepo repository.User, urpService service.UserResourcePermission,tokenService service.Token,otpService service.Otp,emailMediaService service.Media,phoneMediaService service.Media) service.User {
	return &userService{
		userRepo:          userRepo,
		urpService:        urpService,
		tokenService:      tokenService,
		otpService:        otpService,
		emailMediaService: emailMediaService,
		phoneMediaService: phoneMediaService,
	}
}
