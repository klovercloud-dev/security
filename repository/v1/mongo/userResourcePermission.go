package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"time"
)

// UserResourcePermission collection name
var (
	UserResourcePermission = "userResourcePermissionCollection"
)

type userResourcePermissionRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (u userResourcePermissionRepository) Store(userResourcePermission v1.UserResourcePermission) error {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) Get() ([]v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) GetByUserID(userID string) (v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) Delete(userID string) error {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) Update(userResourcePermission v1.UserResourcePermission) error {
	//TODO implement me
	panic("implement me")
}

func NewUserResourcePermissionRepository(m *dmManager, timeout time.Duration) repository.UserResourcePermission {
	return &userResourcePermissionRepository{
		manager: m,
		timeout: timeout,
	}
}
