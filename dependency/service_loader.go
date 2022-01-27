package dependency

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetV1ResourceService() service.Resource {
	var resource service.Resource
	resource = logic.NewResourceService(mongo.NewResourceRepository(3000))
	return resource
}

func GetV1PermissionService() service.Permission {
	var permission service.Permission
	permission = logic.NewPermissionService(mongo.NewPermissionRepository(3000))
	return permission
}
func GetV1HttpClient() service.HttpClient{
	return logic.NewHttpClientService()
}
func GetV1UserService() service.User {
	//if config.RunMode==string(enums.PRODUCTION) {
	//	return logic.NewUserService(mongo.NewUserRepository(3000),GetV1UserResourcePermissionService(),GetV1TokenService(), GetV1OtpService(), GetV1EmailMediaService(), GetV1PhoneMediaService(),GetV1HttpClient())
	//}else if config.RunMode==string(enums.DEVELOP) {
	//	return logic.NewUserMock(logic.NewMockEmailService(), logic.NewMockPhoneService())
	//}else {
	//	return nil
	//}
	return logic.NewUserService(mongo.NewUserRepository(3000),GetV1UserResourcePermissionService(),GetV1TokenService(), GetV1OtpService(), GetV1EmailMediaService(), GetV1PhoneMediaService(),GetV1HttpClient())
}

func GetV1EmailMediaService()service.Media{
	return logic.NewEmailService()
}

func GetV1OtpService()service.Otp{
	//if config.RunMode==string(enums.PRODUCTION)  || config.RunMode==string(enums.DEVELOP){
	//	return logic.NewOtpService(mongo.NewOtpRepository(3000))
	//}
	//return  logic.NewMockOtpService()
	return logic.NewOtpService(mongo.NewOtpRepository(3000))
}

func GetV1PhoneMediaService()service.Media{
	return logic.NewPhoneService()
}
func GetV1JwtService() service.Jwt {
	return  logic.NewJwtService()
}

func GetV1UserResourcePermissionService() service.UserResourcePermission {
	//if config.RunMode==string(enums.PRODUCTION) {
	//	return logic.NewUserResourcePermissionService(mongo.NewUserResourcePermissionRepository(3000))
	//}else if config.RunMode==string(enums.DEVELOP) {
	//	return logic.NewMockUserResourcePermissionService()
	//}else {
	//	return nil
	//}

	//return logic.NewUserResourcePermissionService(mongo.NewUserResourcePermissionRepository(3000))
	return logic.NewUserResourcePermissionService(mongo.NewUserRepository(3000), mongo.NewRoleRepository(3000))
}

func GetV1TokenService() service.Token {
	//if config.RunMode==string(enums.PRODUCTION)  || config.RunMode==string(enums.DEVELOP){
	//	return logic.NewTokenService(mongo.NewTokenRepository(3000), GetV1JwtService())
	//}
	//return logic.NewTokenMock()
	return logic.NewTokenService(mongo.NewTokenRepository(3000), GetV1JwtService())
}

func GetV1RoleService()service.Role{
	//if config.RunMode==string(enums.PRODUCTION)  || config.RunMode==string(enums.DEVELOP){
	//	return logic.NewRoleService(mongo.NewRoleRepository(3000), GetV1PermissionService())
	//}
	//return logic.NewMockRoleService()
	return logic.NewRoleService(mongo.NewRoleRepository(3000), GetV1PermissionService())
}