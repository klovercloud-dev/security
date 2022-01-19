package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci/api/common"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/api"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/enums"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

type roleApi struct {
	service service.Role
	jwtService service.Jwt
}

func (r roleApi) CheckUserPermission(context echo.Context) error {
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
	if !r.jwtService.IsTokenValid(token){
		return common.GenerateForbiddenResponse(context, "[ERROR]: Token is expired!","Please login again to get token!")
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Publickey), nil
	})
	jsonbody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	userResourcePermission := v1.UserResourcePermission{}
	if err := json.Unmarshal(jsonbody, &userResourcePermission); err != nil {
		log.Println(err)
	}
	flag := false
	for _, eachResource := range userResourcePermission.Resources {
		if eachResource.Name == string(enums.USER) {
			for _, eachRole := range eachResource.Roles {
				if eachRole.Name == string(enums.ADMIN) {
					flag = true
				}
			}
		}
	}
	if !flag {
		return common.GenerateForbiddenResponse(context,"[ERROR]: Insufficient permission","User do not have sufficient permission!")
	}
	return nil
}

func (r roleApi) Store(context echo.Context) error {
	if err := r.CheckUserPermission(context); err != nil {
		return nil
	}
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

func NewRoleApi(roleService service.Role, jwtService service.Jwt) api.Role {
	return &roleApi{
		service: roleService,
		jwtService: jwtService,
	}
}