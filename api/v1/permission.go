package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type permissionApi struct {
	service service.Permission
}

func (p permissionApi) Store(context echo.Context) error {
	formData := v1.Permission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := p.service.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func (p permissionApi) Get(context echo.Context) error {
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

func NewPermissionApi(service service.Permission) api.Permission {
	return &permissionApi{service: service}
}
