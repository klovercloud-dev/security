package v1

import (
	"errors"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type roleApi struct {
	service service.Role
	jwtService service.Jwt
}

func (r roleApi) Store(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.CREATE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	formData := v1.Role{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err = r.service.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func (r roleApi) Get(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	data := r.service.Get()
	if len(data) == 0 {
		return common.GenerateErrorResponse(context, nil, "No roles found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r roleApi) GetByName(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	name := context.Param("roleName")
	data := r.service.GetByName(name)
	if data.Name == "" {
		return common.GenerateErrorResponse(context, nil, "Role not found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r roleApi) Delete(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.DELETE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	name := context.Param("roleName")
	err = r.service.Delete(name)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Success!")
}

func (r roleApi) Update(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.UPDATE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	var formData []v1.Permission
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	name := context.Param("roleName")
	roleUpdateOption := v1.RoleUpdateOption{Option: enums.ROLE_UPDATE_OPTION(context.QueryParam("updateOption"))}

	if name == "" {
		log.Println("Role Name Error:", errors.New("empty role name"))
		return common.GenerateErrorResponse(context, nil, "empty role name")
	}
	err = r.service.Update(name, formData, roleUpdateOption)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func NewRoleApi(roleService service.Role, jwtService service.Jwt) api.Role {
	return &roleApi{
		service: roleService,
		jwtService: jwtService,
	}
}