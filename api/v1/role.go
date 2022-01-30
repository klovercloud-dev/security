package v1

import (
	"errors"
	"github.com/klovercloud-ci-cd/security/api/common"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/api"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"github.com/klovercloud-ci-cd/security/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type roleApi struct {
	service    service.Role
	jwtService service.Jwt
}

// Store... Store Api
// @Summary Store api
// @Description Api for storing role
// @Tags Role
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param data body v1.Role true "dto for creating role"
// @Param action path string true "action [create_user] if admin wants to create new user"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/roles [POST]
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

// Get... Get Api
// @Summary Get api
// @Description Api for getting role
// @Tags Role
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} common.ResponseDTO{data=[]v1.Role{}}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/roles [GET]
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

// GetByName... GetByName Api
// @Summary GetByName api
// @Description Api for getting role by name
// @Tags Role
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param name path string false "role name"
// @Success 200 {object} common.ResponseDTO{data=v1.Role{}}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/roles/{name} [GET]
func (r roleApi) GetByName(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	name := context.Param("name")
	data := r.service.GetByName(name)
	if data.Name == "" {
		return common.GenerateErrorResponse(context, nil, "Role not found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

// Delete... Delete Api
// @Summary Delete api
// @Description Api for deleting role by name
// @Tags Role
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param name path string false "role name"
// @Success 200 {object} common.ResponseDTO{}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/roles/{name} [DELETE]
func (r roleApi) Delete(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.DELETE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	name := context.Param("name")
	err = r.service.Delete(name)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Success!")
}

// Update... Update Api
// @Summary Update api
// @Description Api for updating role by name
// @Tags Role
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param name path string false "role name"
// @Success 200 {object} common.ResponseDTO{}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/roles/{name} [PUT]
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
	name := context.Param("name")
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

// NewRoleApi returns api.Role type api
func NewRoleApi(roleService service.Role, jwtService service.Jwt) api.Role {
	return &roleApi{
		service:    roleService,
		jwtService: jwtService,
	}
}
