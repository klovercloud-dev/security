package logic

import (
	"encoding/json"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type UserMock struct {
}

func (u UserMock) GetByEmail(email string) v1.User {
	panic("implement me")
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
	data := `{
				  "id": "6363355f-d35f-4f0a-9696-9364c9a42051",
				  "first_name": "shabrul",
				  "last_name": "islam",
				  "email": "shabrul2451@gmail.com",
				  "password": "bh09743160",
				  "status": "active",
				  "created_date": "2022-01-09 17:59:56.01641726 +0600 +06 m=+3.931900324",
				  "updated_date": "2022-01-09 17:59:56.01641726 +0600 +06 m=+3.931900324",
				  "token": "122sa224as4as1",
				  "refresh_token": "254as45as4d4411sd51a5",
				  "auth_type": "JwtToken"
}`
	user := v1.User{}
	err := json.Unmarshal([]byte(data), &user)
	if err != nil {
		return v1.User{}, err
	}
	if user.ID == id {
		return user, nil
	}
	return v1.User{}, nil
}

func (u UserMock) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserMock() service.User {
	return &UserMock{}
}
