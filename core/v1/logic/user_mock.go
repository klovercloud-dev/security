package logic

import (
	"crypto/rand"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"io"
	"time"
)

type userMock struct {
	emailMediaService service.Media
	phoneMediaService service.Media
}

func (u userMock) AttachCompany(company v1.Company, companyId,token string) error {
	panic("implement me")
}

func (u userMock) GetByOtp(otp string) v1.User {
	panic("implement me")
}

func (u userMock) GetByPhone(phone string) v1.User {
	panic("implement me")
}

func (u userMock) SendOtp(email, phone string) error {
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
		go u.emailMediaService.Listen(otp)
	}else{
		go u.phoneMediaService.Listen(otp)
	}
	return nil
}

func (u userMock) generateOtp(max int) string {
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

func (u userMock) UpdatePassword(user v1.User) error {
	return nil
}

func (u userMock) GetByEmail(email string) v1.User {
	return v1.User{
		Metadata:    v1.UserMetadata{
			CompanyId: "1001",
		},
		ID:          "1",
		FirstName:   "Shahidul",
		LastName:    "islam",
		Email:       "zeromsi.official@gmail.com",
		Phone:       "",
		Password:    "$2a$10$VP2kfzMgzOT.ketk.g4qhOa5Wop3FreHfs8q5x8Flf9dpiX2Gmpze", //1323234
		Status:      "active",
		CreatedDate: time.Now().UTC(),
		UpdatedDate: time.Now().UTC(),
		AuthType:    "password",
	}
}

func (u userMock) UpdateToken(token, refreshToken, existingToken string) error {
	panic("implement me")
}

func (u userMock) Store(user v1.UserRegistrationDto) error {
	//TODO implement me
	panic("implement me")
}

func (u userMock) Get() []v1.User {
	//TODO implement me
	panic("implement me")
}

func (u userMock) GetByID(id string) v1.User{
	return v1.User{
		ID:           "6363355f-d35f-4f0a-9696-9364c9a42051",
		FirstName:    "shabrul",
		LastName:     "islam",
		Email:        "shabrul2451@gmail.com",
		Password:     "$2a$10$VP2kfzMgzOT.ketk.g4qhOa5Wop3FreHfs8q5x8Flf9dpiX2Gmpze", //1323234
		Status:       "active",
		CreatedDate:  time.Now().UTC(),
		UpdatedDate:  time.Now().UTC(),
		AuthType:     "password",
	}
}

func (u userMock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserMock(emailMediaService service.Media,phoneMediaService service.Media) service.User {
	return &userMock{
		emailMediaService: emailMediaService,
		phoneMediaService: phoneMediaService,
	}
}
