package v1

import (
	"errors"
	"github.com/google/uuid"
	"github.com/klovercloud-ci-cd/security/api/common"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/api"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"github.com/klovercloud-ci-cd/security/enums"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type userApi struct {
	userService                   service.User
	userResourcePermissionService service.UserResourcePermission
	otpService                    service.Otp
	jwtService                    service.Jwt
	resourceService               service.Resource
	roleService                   service.Role
}

func (u userApi) UpdateUserResourcePermission(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.UPDATE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	userId := context.Param("id")
	formData := v1.UserResourcePermission{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	formData.UserId = userId
	resourceMap := getResourceMapFromResources(u.resourceService.Get())
	formData = CheckDuplicateData(formData)
	formData.Metadata.CompanyId = userResourcePermission.Metadata.CompanyId
	roleMap := getRoleMapFromRoles(u.roleService.Get())
	formData.Resources = filterOutNonExistingRolesAndResources(roleMap, resourceMap, formData.Resources)
	if err := formData.Validate(); err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Please give valid user resource permission data!")
	}
	err = u.userService.UpdateUserResourcePermissionDto(userId, formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to update!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully updated")
}

// Update... Update Api
// @Summary Update api
// @Description Api for updating users object
// @Tags User
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param action path string true "action type [reset_password/forgot_password/attach_company/update_status]"
// @Param media path string false "media type [users email/phone] if action forgot_password"
// @Param status path string false "status type [inactive/active] if action update_status"
// @Param id path string false "updating users id, if action update_status"
// @Param password_reset_dto body v1.PasswordResetDto true "dto for resetting users password"
// @Param company_dto body v1.Company true "dto for attaching company with user"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/users [PUT]
func (u userApi) Update(context echo.Context) error {
	action := context.QueryParam("action")
	if action == string(enums.RESET_PASSWORD) {
		return u.ResetPassword(context)
	} else if action == string(enums.FORGOT_PASSWORD) {
		return u.ForgotPassword(context)
	} else if action == string(enums.ATTACH_COMPANY) {
		return u.AttachCompany(context)
	} else if action == string(enums.UPDATE_STATUS) {
		return u.UpdateStatus(context)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: No action type is provided!", "Please provide a action type!")
}

func (u userApi) UpdateStatus(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.UPDATE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	status := context.QueryParam("status")
	if enums.STATUS(status) != enums.ACTIVE && enums.STATUS(status) != enums.INACTIVE {
		return common.GenerateErrorResponse(context, "[ERROR]: Invalid update status!", "Please provide a valid update status!")
	}
	userId := context.QueryParam("id")
	user := u.userService.GetByID(userId)
	if user.Metadata.CompanyId != userResourcePermission.Metadata.CompanyId {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Insufficient permission!", "Operation Failed!")
	}
	if user.ID == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User not found!", "Please provide a valid user id!")
	}
	if user.Status == enums.DELETED {
		return common.GenerateErrorResponse(context, "[ERROR]: User not found!", "Please provide a valid user id!")
	}
	err = u.userService.UpdateStatus(userId, enums.STATUS(status))
	if err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Operation Successful!")
}

func (u userApi) AttachCompany(context echo.Context) error {
	bearerToken := context.Request().Header.Get("Authorization")
	if bearerToken == "" {
		return common.GenerateForbiddenResponse(context, "[ERROR]: No token found!", "Please provide a valid token!")
	}
	var token string
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	} else {
		return common.GenerateForbiddenResponse(context, "[ERROR]: No token found!", "Please provide a valid token!")
	}
	if !u.jwtService.IsTokenValid(token) {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Token is expired!", "Please login again to get token!")
	}
	formData := v1.Company{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if formData.Id == "" {
		formData.Id = uuid.New().String()
	}
	err := u.userService.AttachCompany(formData, formData.Id, token)
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to attach company with user", err.Error())
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully attached company with user")
}

func (u userApi) ForgotPassword(context echo.Context) error {
	media := context.QueryParam("media")
	var err error
	if strings.Contains(media, "@") {
		err = u.userService.SendOtp(media, "")
	} else {
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
	if !u.otpService.IsValid(formData.Otp) {
		return common.GenerateErrorResponse(context, "[ERROR]: Invalid Otp", "Please provide a valid otp!")
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

// Registration... Registration Api
// @Summary Registration api
// @Description Api for users registration
// @Tags User
// @Produce json
// @Param Authorization header string true "Insert your access token while adding new user for your company" default(Bearer <Add access token here>)
// @Param data body v1.UserRegistrationDto true "dto for creating user"
// @Param action path string true "action [create_user] if admin wants to create new user"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/users [POST]
func (u userApi) Registration(context echo.Context) error {
	registrationType := context.QueryParam("action")
	if registrationType == "" {
		return u.registerAdmin(context)
	} else if registrationType == string(enums.CREATE_USER) {
		return u.registerUser(context)
	}
	return common.GenerateErrorResponse(context, "[ERROR]: Failed to register user!", errors.New("invalid query action").Error())
}

func (u userApi) registerAdmin(context echo.Context) error {
	formData := v1.UserRegistrationDto{}
	if err := context.Bind(&formData); err != nil {
		log.Println("Input Error:", err.Error())
		return common.GenerateErrorResponse(context, nil, "Failed to Bind Input!")
	}
	if formData.Password == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to register user!", "password is required")
	} else if len(formData.Password) < 8 {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to register user!", "password length must be at least 8")
	}
	formData.ID = uuid.New().String()
	userResourcePermissionDto := v1.UserResourcePermission{
		Metadata: v1.UserMetadata{},
		UserId:   formData.ID,
	}
	var resourceWiseRoles []v1.ResourceWiseRolesDto
	existingResources := u.resourceService.Get()
	adminRole := u.roleService.GetByName(string(enums.ADMIN))
	for _, each := range existingResources {
		resourceWiseRole := v1.ResourceWiseRolesDto{
			Name:  each.Name,
			Roles: []v1.RoleDto{{Name: adminRole.Name}},
		}
		resourceWiseRoles = append(resourceWiseRoles, resourceWiseRole)
	}
	userResourcePermissionDto.Resources = resourceWiseRoles
	formData.CreatedDate = time.Now().UTC()
	formData.UpdatedDate = time.Now().UTC()
	formData.Status = enums.ACTIVE
	formData.ResourcePermission = userResourcePermissionDto
	err := formData.Validate()
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to register user!", err.Error())
	}
	err = u.userService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	return common.GenerateSuccessResponse(context, formData, nil, "Successfully Created User!")
}

func (u userApi) registerUser(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), string(enums.ADMIN), ""); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	if userResourcePermission.Metadata.CompanyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User got no company!", "Please attach a company first!")
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
	formData.ResourcePermission.Metadata.CompanyId = userResourcePermission.Metadata.CompanyId
	formData.Metadata.CompanyId = userResourcePermission.Metadata.CompanyId
	formData.CreatedDate = time.Now().UTC()
	formData.UpdatedDate = time.Now().UTC()
	formData.Status = enums.ACTIVE
	formData.Metadata = userResourcePermission.Metadata
	formData.ResourcePermission = CheckDuplicateData(formData.ResourcePermission)
	roleMap := getRoleMapFromRoles(u.roleService.Get())
	resourceMap := getResourceMapFromResources(u.resourceService.Get())
	formData.ResourcePermission.Resources = filterOutNonExistingRolesAndResources(roleMap, resourceMap, formData.ResourcePermission.Resources)
	err = formData.Validate()
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to register user!", err.Error())
	}
	err = u.userService.Store(formData)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, err.Error())
	}
	err = u.userService.SendOtp(formData.Email, "")
	if err != nil {
		return common.GenerateErrorResponse(context, "[ERROR]: Failed to send otp!", "User has been created but failed to send otp!")
	}
	return common.GenerateSuccessResponse(context, formData, nil, "Successfully Created User!")
}

// Get... Get Api
// @Summary Get api
// @Description Api for getiing all user by admins company Id
// @Tags User
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param status path string true "status type [active/inactive]"
// @Success 200 {object} common.ResponseDTO{data=[]v1.User{}}
// @Forbidden 403 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Router /api/v1/users [GET]
func (u userApi) Get(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.READ)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	companyId := userResourcePermission.Metadata.CompanyId
	if companyId == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User got no company!", "Please attach a company first!")
	}
	status := context.QueryParam("status")
	if status == string(enums.ACTIVE) {
		return common.GenerateSuccessResponse(context, u.userService.GetUsersByCompanyId(companyId, enums.STATUS(status)), nil, "Success!")
	} else if status == string(enums.INACTIVE) {
		return common.GenerateSuccessResponse(context, u.userService.GetUsersByCompanyId(companyId, enums.STATUS(status)), nil, "Success!")
	}
	return common.GenerateForbiddenResponse(context, "[ERROR]: No valid status found!", "Please provide a valid status.")
}

// GetByID... GetByID Api
// @Summary Registration api
// @Description Api for getiing user by id
// @Tags User
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id user id"
// @Success 200 {object} common.ResponseDTO{data=v1.User{}}
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/users/{id} [GET]
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
	data := u.userService.GetByID(id)
	if data.ID == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User Not Found!", "Please give a valid user id!")
	}
	if data.Metadata.CompanyId != userResourcePermission.Metadata.CompanyId {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Insufficient permission!", "Operation Failed!")
	}
	return common.GenerateSuccessResponse(context, data, nil, "Success!")
}

// Delete... Delete Api
// @Summary Delete api
// @Description Api to delete user
// @Tags User
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param id path string true "id user id"
// @Success 200 {object} common.ResponseDTO
// @Failure 400 {object} common.ResponseDTO
// @Forbidden 403 {object} common.ResponseDTO
// @Router /api/v1/users [DELETE]
func (u userApi) Delete(context echo.Context) error {
	userResourcePermission, err := GetUserResourcePermissionFromBearerToken(context, u.jwtService)
	if err != nil {
		return common.GenerateErrorResponse(context, err.Error(), "Operation Failed!")
	}
	if err := checkAuthority(userResourcePermission, string(enums.USER), "", string(enums.DELETE)); err != nil {
		return common.GenerateForbiddenResponse(context, err.Error(), "Operation Failed!")
	}
	id := context.Param("id")
	user := u.userService.GetByID(id)
	if user.Metadata.CompanyId != userResourcePermission.Metadata.CompanyId {
		return common.GenerateForbiddenResponse(context, "[ERROR]: Insufficient permission!", "Operation Failed!")
	}
	if user.ID == "" {
		return common.GenerateErrorResponse(context, "[ERROR]: User not found!", "Please provide a valid user id!")
	}
	err = u.userService.Delete(id)
	if err != nil {
		return common.GenerateErrorResponse(context, nil, "Failed to Delete User!")
	}
	return common.GenerateSuccessResponse(context, nil, nil, "Successfully Deleted User!")
}

// NewUserApi returns api.User type api
func NewUserApi(userService service.User, userResourcePermissionService service.UserResourcePermission, otpService service.Otp, jwtService service.Jwt, resourceService service.Resource, roleService service.Role) api.User {
	return &userApi{
		userService:                   userService,
		userResourcePermissionService: userResourcePermissionService,
		otpService:                    otpService,
		jwtService:                    jwtService,
		resourceService:               resourceService,
		roleService:                   roleService,
	}
}
