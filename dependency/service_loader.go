package dependency

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetV1ResourceService() service.Resource {
	var resource service.Resource
	resource = logic.NewResourceService(mongo.NewResourceRepository(300))

	return resource
}

func GetV1PermissionService() service.Permission {
	var permission service.Permission
	permission = logic.NewPermissionService(mongo.NewPermissionRepository(300))

	return permission
}

func GetV1UserService() service.User {
	var user service.User
	user = logic.NewUserService(mongo.NewUserRepository(300), mongo.NewUserResourcePermissionRepository(300))

	return user
}
