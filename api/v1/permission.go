package v1

import (
	"github.com/klovercloud-ci-cd/security/api/common"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/api"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"github.com/klovercloud-ci-cd/security/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type permissionApi struct {
	service    service.Permission
	jwtService service.Jwt
}

func (p permissionApi) Store(context echo.Context) error {
	formData := v1.Permission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	err := p.service.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

// Get... Get Api
// @Summary Get api
// @Description Api for getting permissions
// @Tags Permission
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} common.ResponseDTO{data=[]v1.Permission{}}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/permissions [GET]
func (p permissionApi) Get(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, p.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	data := p.service.Get()
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (p permissionApi) Delete(context echo.Context) error {
	name := context.QueryParam("permissionName")
	err := p.service.Delete(name)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Success!")
}

// NewPermissionApi returns api.Permission type api
func NewPermissionApi(service service.Permission, jwtService service.Jwt) api.Permission {
	return &permissionApi{
		service:    service,
		jwtService: jwtService,
	}
}
