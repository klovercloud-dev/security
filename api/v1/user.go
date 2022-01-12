package v1

import (
	"github.com/google/uuid"
	"github.com/klovercloud-ci/api/common"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type userApi struct {
	userService service.User
	userResourcePermissionService service.UserResourcePermission
	otpService service.Otp
}

func (u userApi) UpdateUserResourcePermission(context echo.Context) error {
	userId:=context.Param("id")
	formData := v1.UserResourcePermission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	formData.UserId=userId
	err := u.userResourcePermissionService.Update(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to update!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully updated")
}

func (u userApi) Update(context echo.Context) error {
	action := context.QueryParam("action")
	if action == string(enums.RESET_PASSWORD) {
		return u.ResetPassword(context)
	} else if action == string(enums.FORGOT_PASSWORD) {
		return u.ForgotPassword(context)
	} else if action == string(enums.ATTACH_COMPANY) {
		return u.AttachCompany(context)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: No action type is provided!", "Please provide a action type!")
}

func (u userApi) AttachCompany(context echo.Context) error {
	return nil
}
func (u userApi) ForgotPassword(context echo.Context) error {
	media := context.QueryParam("media")
	var err error
	if strings.Contains(media,"@") {
		err = u.userService.SendOtp(media, "")
	} else  {
		err = u.userService.SendOtp("", media)
	}
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to generate OTP", err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Please check your corresponding media to get the otp")
}
func (u userApi) ResetPassword(context echo.Context) error {
	formData := v1.PasswordResetDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}

	if !u.otpService.IsValid(formData.Otp){
		return common.GenerateErrorResponse(context,"[ERROR]: Invalid Otp","Please provide a valid otp!")
	}
	var user v1.User
	if formData.Otp!=""{
		user=u.userService.GetByID(formData.Otp)
	}else {
		user = u.userService.GetByEmail(formData.Email)
		if user.ID == "" {
			return common.GenerateForbiddenResponse(context, "[ERROR]: No User found!", "Please login with actual user email!")
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.CurrentPassword))
		if err != nil {
			return common.GenerateForbiddenResponse(context, "[ERROR]: Password not matched!", "Please provide due credential!"+err.Error())
		}
	}
	user.Password = formData.NewPassword
	err := u.userService.UpdatePassword(user)
	if err != nil {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Failed to reset password!", err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Operation Successful!")
}

func (u userApi) UserResourcePermissionApi(context echo.Context) error {
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

func (u userApi) Store(context echo.Context) error {
	formData := v1.UserRegistrationDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	formData.ID = uuid.New().String()
	formData.CreatedDate = time.Now().UTC()
	formData.UpdatedDate = time.Now().UTC()
	formData.Status=enums.ACTIVE
	err:=formData.Validate()
	if err!=nil{
		return common.GenerateErrorResponse(context,"[ERROR]: Failed to register user!",err.Error())
	}
	err = u.userService.Store(formData)
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
	data:= u.userService.GetByID(id)
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

func NewUserApi(userService service.User,userResourcePermissionService service.UserResourcePermission,	otpService service.Otp) api.User {
	return &userApi{
		userService: userService,
		userResourcePermissionService: userResourcePermissionService,
		otpService: otpService,
	}
}
