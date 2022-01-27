package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type userResourcePermissionService struct {
	userRepo repository.User
	roleRepo repository.Role
}

func (u userResourcePermissionService) GetByUserID(userID string) v1.UserResourcePermissionDto {
	user := u.userRepo.GetByID(userID)
	userResourcePermission := v1.UserResourcePermissionDto{
		Metadata:  user.Metadata,
		UserId:    userID,
	}
	resourceWiseRoles := []v1.ResourceWiseRoles{}
	for _, eachResource := range user.ResourcePermission.Resources {
		resourceWiseRole := v1.ResourceWiseRoles{Name:  eachResource.Name}
		for _, eachRole := range eachResource.Roles {
			resourceWiseRole.Roles = append(resourceWiseRole.Roles,  u.roleRepo.GetByName(eachRole.Name))
		}
		resourceWiseRoles = append(resourceWiseRoles, resourceWiseRole)
	}
	userResourcePermission.Resources = resourceWiseRoles
	return userResourcePermission
}

func NewUserResourcePermissionService(userRepo repository.User, roleRepo repository.Role) service.UserResourcePermission {
	return &userResourcePermissionService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
