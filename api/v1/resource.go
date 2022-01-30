package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
)

type resourceApi struct {
	service service.Resource
	jwtService service.Jwt
}

func (r resourceApi) Store(context echo.Context) error {
	formData := v1.Resource{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := r.service.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

// Get... Get Api
// @Summary Store api
// @Description Api for getting resources
// @Tags Resource
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} common.ResponseDTO{data=[]v1.Resource{}}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/resources [GET]
func (r resourceApi) Get(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, r.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	data := r.service.Get()
	if len(data) == 0 {
		return common.GenerateErrorResponse(context, nil, "No resource found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r resourceApi) GetByName(context echo.Context) error {
	name := context.Param("resourceName")
	data := r.service.GetByName(name)
	if data.Name == "" {
		return common.GenerateErrorResponse(context, nil, "Resource not found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r resourceApi) Delete(context echo.Context) error {
	name := context.Param("resourceName")
	err := r.service.Delete(name)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Success!")
}

func NewResourceApi(resourceService service.Resource, jwtService service.Jwt) api.Resource {
	return &resourceApi{
		service: resourceService,
		jwtService: jwtService,
	}
}
