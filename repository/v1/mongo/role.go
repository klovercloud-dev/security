package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"time"
)

// RoleCollection collection name
var (
	RoleCollection = "roleCollection"
)

type roleRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (r roleRepository) Store(role v1.Role) error {
	//TODO implement me
	panic("implement me")
}

func (r roleRepository) Get() ([]v1.Role, int64) {
	//TODO implement me
	panic("implement me")
}

func (r roleRepository) GetByName(name string) (v1.Role, error) {
	//TODO implement me
	panic("implement me")
}

func (r roleRepository) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

func (r roleRepository) AppendPermissions(name string, permissions []v1.Permission) error {
	//TODO implement me
	panic("implement me")
}

func (r roleRepository) RemovePermissions(name string, permission []v1.Permission) error {
	//TODO implement me
	panic("implement me")
}

func NewRoleRepository(manager *dmManager, timeout time.Duration) repository.Role {
	return &roleRepository{manager: manager, timeout: timeout}
}
