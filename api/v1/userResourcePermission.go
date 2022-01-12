package v1

import (
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
)

type userResourcePermissionApi struct {
	userResourcePermissionService service.UserResourcePermission
}

func (u userResourcePermissionApi) Store(context echo.Context) error {
	formData := v1.UserResourcePermission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := u.userResourcePermissionService.Store(formData)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func (u userResourcePermissionApi) Get(context echo.Context) error {
	data := u.userResourcePermissionService.Get()
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (u userResourcePermissionApi) GetByUserID(context echo.Context) error {
	id := context.Param("id")
	data := u.userResourcePermissionService.GetByUserID(id)

	if data.UserId == "" {
		return common.GenerateErrorResponse(context, nil, "User Not Found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (u userResourcePermissionApi) Delete(context echo.Context) error {
	id := context.Param("id")
	err := u.userResourcePermissionService.Delete(id)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to Delete User!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully Deleted User!")
}

func (u userResourcePermissionApi) Update(context echo.Context) error {
	formData := v1.UserResourcePermission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := u.userResourcePermissionService.Update(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to update!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully updated")
}

func NewUserResourcePermissionApi(userResourcePermissionService service.UserResourcePermission) api.UserResourcePermission {
	return &userResourcePermissionApi{
		userResourcePermissionService: userResourcePermissionService,
	}
}