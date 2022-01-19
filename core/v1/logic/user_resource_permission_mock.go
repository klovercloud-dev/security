package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type mockUserResourcePermissionService struct {
}
 
func (m mockUserResourcePermissionService) Store(userResourcePermission v1.UserResourcePermission) error {
	panic("implement me")
}

func (m mockUserResourcePermissionService) Get() []v1.UserResourcePermission {
	panic("implement me")
}

func (m mockUserResourcePermissionService) GetByUserID(userID string) v1.UserResourcePermission {
	return v1.UserResourcePermission{
		Metadata: v1.UserMetadata{CompanyId: "1"},
		UserId:    "1",
		Resources: []v1.ResourceWiseRoles{
			{
				Name: "pipeline",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name:"CREATE"},{Name: "READ"},{Name: "UPDATE"},{Name: "DELETE"}}}},
			},
			{
				Name: "company",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name:"CREATE"},{Name: "READ"},{Name: "UPDATE"},{Name: "DELETE"}}}},
			},
		},
	}
}

func (m mockUserResourcePermissionService) Delete(userID string) error {
	panic("implement me")
}

func (m mockUserResourcePermissionService) Update(userResourcePermission v1.UserResourcePermission) error {
	panic("implement me")
}

func NewMockUserResourcePermissionService() service.UserResourcePermission {
	return &mockUserResourcePermissionService{
	}
}
