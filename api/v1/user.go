package v1

import (
	"github.com/google/uuid"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

type userApi struct {
	userService service.User
}

func (u userApi) Store(context echo.Context) error {
	formData := v1.UserRegistrationDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	formData.ID = uuid.New().String()
	formData.CreatedDate = time.Now().String()
	err := u.userService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, formData, nil, "Successfully Created User!")
}

func (u userApi) Get(context echo.Context) error {
	data := u.userService.Get()
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (u userApi) GetByID(context echo.Context) error {
	id := context.Param("id")
	data, _ := u.userService.GetByID(id)

	if data.ID == "" {
		return common.GenerateErrorResponse(context, nil, "User Not Found!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

func (u userApi) Delete(context echo.Context) error {
	id := context.Param("id")
	err := u.userService.Delete(id)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to Delete User!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully Deleted User!")
}

func NewUserApi(userService service.User) api.User {
	return &userApi{
		userService: userService,
	}
}
