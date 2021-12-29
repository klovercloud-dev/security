package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type resourceApi struct {
	resourceService service.Resource
}

func (r resourceApi) Store(context echo.Context) error {
	formData := v1.Resource{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := r.resourceService.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func (r resourceApi) Get(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r resourceApi) Delete(context echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewResourceApi(resourceService service.Resource) api.Resource {
	return &resourceApi{resourceService: resourceService}
}
