package logic

import (
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
)

type mockRoleService struct {
}

func (m mockRoleService) Store(role v1.RoleDto) error {
	panic("implement me")
}

func (m mockRoleService) Get() []v1.RoleDto {
	panic("implement me")
}

func (m mockRoleService) GetByName(name string) v1.RoleDto {
	panic("implement me")
}

func (m mockRoleService) Delete(name string) error {
	panic("implement me")
}

func (m mockRoleService) Update(name string, permissions []v1.Permission, option v1.RoleUpdateOption) error {
	panic("implement me")
}

// NewMockRoleService returns service.RoleDto type service
func NewMockRoleService() service.Role {
	return &mockRoleService{}
}
