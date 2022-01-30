package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
)

type mockUserResourcePermissionService struct {
}

var mockUserResourcePermissions map[string]v1.UserResourcePermissionDto

//InitMockUserResourcePermissions init mock data
func InitMockUserResourcePermissions() {
	mockUserResourcePermissions = make(map[string]v1.UserResourcePermissionDto)
	mockUserResourcePermissions["b876ec8a-9650-408e-84bb-e5f3d36b4704"] = v1.UserResourcePermissionDto{
		Metadata: v1.UserMetadata{CompanyId: "12345"},
		UserId:   "b876ec8a-9650-408e-84bb-e5f3d36b4704",
		Resources: []v1.ResourceWiseRoles{
			{
				Name:  "pipeline",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
			{
				Name:  "process",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
			{
				Name:  "company",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
			{
				Name:  "repository",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
			{
				Name:  "application",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
			{
				Name:  "user",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
		},
	}

	mockUserResourcePermissions["9572c6dd-96a0-4e40-a01e-56bf1f7d3c59"] = v1.UserResourcePermissionDto{
		Metadata: v1.UserMetadata{CompanyId: ""},
		UserId:   "9572c6dd-96a0-4e40-a01e-56bf1f7d3c59",
		Resources: []v1.ResourceWiseRoles{
			{
				Name:  "pipeline",
				Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name: "CREATE"}, {Name: "READ"}, {Name: "UPDATE"}, {Name: "DELETE"}}}},
			},
		},
	}
}

func (m mockUserResourcePermissionService) Store(userResourcePermission v1.UserResourcePermissionDto) error {
	panic("implement me")
}

func (m mockUserResourcePermissionService) Get() []v1.UserResourcePermissionDto {
	panic("implement me")
}

func (m mockUserResourcePermissionService) GetByUserID(userID string) v1.UserResourcePermissionDto {
	InitMockUserResourcePermissions()
	if userResourcePermission, ok := mockUserResourcePermissions[userID]; ok {
		return userResourcePermission
	}
	return v1.UserResourcePermissionDto{}
}

func (m mockUserResourcePermissionService) Delete(userID string) error {
	panic("implement me")
}

func (m mockUserResourcePermissionService) Update(userResourcePermission v1.UserResourcePermissionDto) error {
	panic("implement me")
}

// NewMockUserResourcePermissionService returns service.UserResourcePermission type service
func NewMockUserResourcePermissionService() service.UserResourcePermission {
	return &mockUserResourcePermissionService{}
}
