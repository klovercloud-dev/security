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
}

func (r roleApi) Store(context echo.Context) error {
	formData := v1.Role{}
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

func (r roleApi) Get(context echo.Context) error {
	data := r.service.Get()
	if len(data) == 0 {
		return common.GenerateErrorResponse(context, nil, "No roles found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r roleApi) GetByName(context echo.Context) error {
	name := context.Param("roleName")
	data := r.service.GetByName(name)
	if data.Name == "" {
		return common.GenerateErrorResponse(context, nil, "Role not found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (r roleApi) Delete(context echo.Context) error {
	name := context.Param("roleName")
	err := r.service.Delete(name)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Success!")
}

func (r roleApi) Update(context echo.Context) error {
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
	err := r.service.Update(name, formData, roleUpdateOption)
	if err != nil {
		log.Println("[Error]:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, formData,
		nil, "Operation Successful")
}

func NewRoleApi(roleService service.Role) api.Role {
	return &roleApi{service: roleService}
}