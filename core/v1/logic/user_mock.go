package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"time"
)

type UserMock struct {
}

func (u UserMock) UpdatePassword(user v1.User) error {
	return nil
}

func (u UserMock) GetByEmail(email string) v1.User {
	return v1.User{
		ID:           "1",
		FirstName:    "Shahidul",
		LastName:     "islam",
		Email:        "shahidul.islam@gmail.com",
		Password:     "$2a$10$VP2kfzMgzOT.ketk.g4qhOa5Wop3FreHfs8q5x8Flf9dpiX2Gmpze", //1323234
		Status:       "active",
		CreatedDate:  time.Now().UTC(),
		UpdatedDate:  time.Now().UTC(),
		AuthType:     "password",
	}
}

func (u UserMock) UpdateToken(token, refreshToken, existingToken string) error {
	panic("implement me")
}

func (u UserMock) Store(user v1.UserRegistrationDto) error {
	//TODO implement me
	panic("implement me")
}

func (u UserMock) Get() []v1.User {
	//TODO implement me
	panic("implement me")
}

func (u UserMock) GetByID(id string) (v1.User, error) {
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
	}, nil
}

func (u UserMock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserMock() service.User {
	return &UserMock{}
}
