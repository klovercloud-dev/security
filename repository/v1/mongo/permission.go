package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"time"
)

// PermissionCollection collection name
var (
	PermissionCollection = "permissionCollection"
)

type permissionRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p permissionRepository) Store(permission v1.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (p permissionRepository) Get() []v1.Permission {
	//TODO implement me
	panic("implement me")
}

func (p permissionRepository) Delete(permissionName string) error {
	//TODO implement me
	panic("implement me")
}

func NewPermissionRepository(timeout int) repository.Permission {
	return &permissionRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout) * time.Second,
	}
}
