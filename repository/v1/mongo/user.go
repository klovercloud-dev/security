package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"time"
)

// UserCollection collection name
var (
	UserCollection = "userCollection"
)

type userRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (u userRepository) Store(user v1.User) error {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Get() ([]v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByID(id string) (v1.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(manager *dmManager, timeout time.Duration) repository.User {
	return &userRepository{
		manager: manager,
		timeout: timeout,
	}
}
