package v1

import (
	"errors"
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
	jwtService service.Jwt
	resourceService service.Resource
	roleService service.Role
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
	} else if action == string(enums.ACTIVE) || action == string(enums.INACTIVE) {
		return u.UpdateStatus(context)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: No action type is provided!", "Please provide a action type!")
}

func (u userApi) UpdateStatus(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), string(enums.ADMIN), ""); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	action := context.QueryParam("action")
	userId := context.QueryParam("id")
	user := u.userService.GetByID(userId)
	if user.ID == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User not found!", "Please provide a valid user id!")
	}
	err = u.userService.UpdateStatus(userId, enums.STATUS(action))
	if err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Operation Successful!")
}

func (u userApi) AttachCompany(context echo.Context) error {
	bearerToken:=context.Request().Header.Get("Authorization")
	if bearerToken==""{
		return common.GenerateForbiddenResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
	}
	var token string
	if len(strings.Split(bearerToken," "))==2{
		token=strings.Split(bearerToken," ")[1]
	}else{
		return common.GenerateForbiddenResponse(context,"[ERROR]: No token found!","Please provide a valid token!")
	}
	if !u.jwtService.IsTokenValid(token){
		return common.GenerateForbiddenResponse(context, "[ERROR]: Token is expired!","Please login again to get token!")
	}
	formData := v1.Company{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	err := u.userService.AttachCompany(formData, formData.Id,token)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to attach company with user", err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully attached company with user")
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

	user = u.userService.GetByEmail(formData.Email)
	if user.ID == "" {
		return common.GenerateForbiddenResponse(context, "[ERROR]: No User found!", "Please login with actual user email!")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.CurrentPassword))
	if err != nil {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Password not matched!", "Please provide due credential!"+err.Error())
	}

	user.Password = formData.NewPassword
	err = u.userService.UpdatePassword(user)
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

func (u userApi) Registration(context echo.Context) error {
	registrationType := context.QueryParam("action")
	if registrationType == "" {
		return u.RegisterAdmin(context)
	} else if registrationType == string(enums.CREATE_USER) {
		return u.RegisterUser(context)
	}
	return common.GenerateErrorResponse(context,"[ERROR]: Failed to register user!",errors.New("invalid query action").Error())
}

func (u userApi) RegisterAdmin(context echo.Context) error {
	formData := v1.UserRegistrationDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if formData.Password == "" {
		return common.GenerateErrorResponse(context,"[ERROR]: Failed to register user!", "password is required")
	} else if len(formData.Password) < 8 {
		return common.GenerateErrorResponse(context,"[ERROR]: Failed to register user!", "password length must be at least 8")
	}
	formData.ID = uuid.New().String()
	userResourcePermission:=v1.UserResourcePermission{
		Metadata:  v1.UserMetadata{},
		UserId:   formData.ID ,
	}
	var resourceWiseRoles []v1.ResourceWiseRoles
	existingResources := u.resourceService.Get()
	roles := u.roleService.GetByName(string(enums.ADMIN))
	for _, each := range existingResources {
		resourceWiseRole := v1.ResourceWiseRoles{
			Name:  each.Name,
			Roles: []v1.Role{roles},
		}
		resourceWiseRoles = append(resourceWiseRoles, resourceWiseRole)
	}
	userResourcePermission.Resources = resourceWiseRoles
	formData.ResourcePermission = userResourcePermission

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

func (u userApi) RegisterUser(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), string(enums.ADMIN), ""); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	if userResourcePermission.Metadata.CompanyId==""{
		return common.GenerateErrorResponse(context,"[ERROR]: User got no company!","Please attach a company first!")
	}
	formData := v1.UserRegistrationDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if formData.Password != "" {
		formData.Password = ""
	}
	formData.ID = uuid.New().String()
	var resourceWiseRoles []v1.ResourceWiseRoles
	resources := formData.ResourcePermission.Resources
	existingResources := u.resourceService.Get()
	existingRoles := u.roleService.Get()
	roleMap:=make(map[string]v1.Role)
	resourceMap:=make(map[string]v1.Resource)
	for _,role:=range existingRoles{
		roleMap[role.Name]=role
	}
	existingRoles=nil
	for _,resource:=range existingResources{
		resourceMap[resource.Name]=resource
	}
	existingResources=nil
	for _, eachResource := range resources {
		if _, ok := resourceMap[eachResource.Name]; ok {
			var addedRoles []v1.Role
			for _, eachRole := range eachResource.Roles {
				if val, roleOk := roleMap[eachRole.Name]; roleOk {
					addedRoles = append(addedRoles, val)
				}
			}
			resourceWiseRole := v1.ResourceWiseRoles{
				Name:  eachResource.Name,
				Roles: addedRoles,
			}
			resourceWiseRoles = append(resourceWiseRoles, resourceWiseRole)
		}
	}

	userResourcePermission.Resources = resourceWiseRoles
	formData.ResourcePermission = userResourcePermission
	formData.CreatedDate = time.Now().UTC()
	formData.UpdatedDate = time.Now().UTC()
	formData.Status=enums.ACTIVE
	formData.Metadata=userResourcePermission.Metadata
	err = formData.Validate()
	if err!=nil{
		return common.GenerateErrorResponse(context,"[ERROR]: Failed to register user!",err.Error())
	}
	err = u.userService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	err=u.userService.SendOtp(formData.Email,"")
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to send otp!", "User has been created but failed to send otp!")
	}
	return common.GenerateSuccessResponse(context, formData, nil, "Successfully Created User!")
}

func (u userApi) Get(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), string(enums.ADMIN), ""); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	companyId := userResourcePermission.Metadata.CompanyId
	if companyId == ""{
		return common.GenerateErrorResponse(context,"[ERROR]: User got no company!","Please attach a company first!")
	}
	return common.GenerateSuccessResponse(context, u.userService.GetUsersByCompanyId(companyId), nil, "Success!")
}

func (u userApi) GetByID(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	id := context.Param("id")
	if userResourcePermission.UserId != id {
		if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
			return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
		}
	}
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

func NewUserApi(userService service.User,userResourcePermissionService service.UserResourcePermission, otpService service.Otp, jwtService service.Jwt, resourceService service.Resource, roleService service.Role) api.User {
	return &userApi{
		userService: userService,
		userResourcePermissionService: userResourcePermissionService,
		otpService: otpService,
		jwtService: jwtService,
		resourceService: resourceService,
		roleService: roleService,
	}
}