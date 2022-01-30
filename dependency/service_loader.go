package dependency

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

// GetV1ResourceService returns service.Resource
func GetV1ResourceService() service.Resource {
	var resource service.Resource
	resource = logic.NewResourceService(mongo.NewResourceRepository(3000))
	return resource
}

// GetV1PermissionService returns service.Permission
func GetV1PermissionService() service.Permission {
	var permission service.Permission
	permission = logic.NewPermissionService(mongo.NewPermissionRepository(3000))
	return permission
}

// GetV1HttpClient returns service.HttpClient
func GetV1HttpClient() service.HttpClient {
	return logic.NewHttpClientService()
}

// GetV1UserService returns service.User
func GetV1UserService() service.User {
	return logic.NewUserService(mongo.NewUserRepository(3000), GetV1UserResourcePermissionService(), GetV1TokenService(), GetV1OtpService(), GetV1EmailMediaService(), GetV1PhoneMediaService(), GetV1HttpClient())
}

// GetV1EmailMediaService returns service.Media
func GetV1EmailMediaService() service.Media {
	return logic.NewEmailService()
}

// GetV1OtpService returns service.Otp
func GetV1OtpService() service.Otp {
	return logic.NewOtpService(mongo.NewOtpRepository(3000))
}

// GetV1PhoneMediaService returns service.Media
func GetV1PhoneMediaService() service.Media {
	return logic.NewPhoneService()
}

// GetV1JwtService returns service.Jwt
func GetV1JwtService() service.Jwt {
	return logic.NewJwtService()
}

// GetV1UserResourcePermissionService returns service.UserResourcePermission
func GetV1UserResourcePermissionService() service.UserResourcePermission {
	return logic.NewUserResourcePermissionService(mongo.NewUserRepository(3000), mongo.NewRoleRepository(3000))
}

// GetV1TokenService returns service.Token
func GetV1TokenService() service.Token {
	return logic.NewTokenService(mongo.NewTokenRepository(3000), GetV1JwtService())
}

// GetV1RoleService returns service.Role
func GetV1RoleService() service.Role {
	return logic.NewRoleService(mongo.NewRoleRepository(3000), GetV1PermissionService())
}
