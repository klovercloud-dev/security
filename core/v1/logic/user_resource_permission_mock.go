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

func (m mockUserResourcePermissionService) Get() ([]v1.UserResourcePermission, error) {
	panic("implement me")
}

func (m mockUserResourcePermissionService) GetByUserID(userID string) (v1.UserResourcePermission, error) {
	return v1.UserResourcePermission{
		UserId:    "1",
		Resources: []struct {
			Name  string    `json:"name" bson:"name"`
			Roles []v1.Role `json:"roles" bson:"roles"`
		}{
			{Name: "Pipeline", Roles: []v1.Role{{Name: "ADMIN", Permissions: []v1.Permission{{Name:"CREATE"},{Name: "READ"},{Name: "UPDATE"},{Name: "DELETE"}}}}},

		},
	},nil
}

func (m mockUserResourcePermissionService) Delete(userID string) error {
	panic("implement me")
}

func (m mockUserResourcePermissionService) Update(userResourcePermission v1.UserResourcePermission) error {
	panic("implement me")
}

func NewMockUserResourcePermissionService() service.UserResourcePermission {
	//return &mockUserResourcePermissionService{
	//}
	return nil
}
