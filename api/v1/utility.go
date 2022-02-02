package v1

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/klovercloud-ci-cd/security/config"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/service"
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

// GetUserResourcePermissionFromBearerToken returns users resource wise permissions from bearer token
func GetUserResourcePermissionFromBearerToken(context echo.Context, jwtService service.Jwt) (v1.UserResourcePermissionDto, error) {
	bearerToken := context.Request().Header.Get("Authorization")
	if bearerToken == "" {
		return v1.UserResourcePermissionDto{}, errors.New("[ERROR]: No token found!")
	}
	var token string
	if len(strings.Split(bearerToken, " ")) == 2 {
		token = strings.Split(bearerToken, " ")[1]
	} else {
		return v1.UserResourcePermissionDto{}, errors.New("[ERROR]: No token found!")
	}
	if !jwtService.IsTokenValid(token) {
		return v1.UserResourcePermissionDto{}, errors.New("[ERROR]: Token is expired!")
	}
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Publickey), nil
	})
	jsonbody, err := json.Marshal(claims["data"])
	if err != nil {
		log.Println(err)
	}
	userResourcePermission := v1.UserResourcePermissionDto{}
	if err := json.Unmarshal(jsonbody, &userResourcePermission); err != nil {
		return v1.UserResourcePermissionDto{}, errors.New("[ERROR]: No resource permissions!")
	}
	return userResourcePermission, nil
}

func checkAuthority(userResourcePermission v1.UserResourcePermissionDto, resourceName, role, permission string) error {
	var resourceWiseRoles v1.ResourceWiseRolesDto
	for _, resource := range userResourcePermission.Resources {
		if resource.Name == resourceName {
			resourceWiseRoles = resource
			break
		}
	}
	if role != "" {
		for _, each := range resourceWiseRoles.Roles {
			if each.Name == role {
				return nil
			}
		}
	} else if permission != "" {
		for _, each := range resourceWiseRoles.Roles {
			for _, perm := range each.Permissions {
				if perm.Name == permission {
					return nil
				}
			}

		}
	}
	return errors.New("[ERROR]: Insufficient permission")
}

func getRoleMapFromRoles(roles []v1.RoleDto) map[string]v1.Role {
	roleMap := make(map[string]v1.Role)
	for _, role := range roles {
		roleMap[role.Name] = v1.Role{Name: role.Name}
	}
	return roleMap
}

func getResourceMapFromResources(resources []v1.Resource) map[string]v1.Resource {
	resourceMap := make(map[string]v1.Resource)
	for _, resource := range resources {
		resourceMap[resource.Name] = resource
	}
	return resourceMap
}

func filterOutNonExistingRolesAndResources(roleMap map[string]v1.Role, resourceMap map[string]v1.Resource, resourceWiseRoles []v1.ResourceWiseRoles) []v1.ResourceWiseRoles {
	var newResourceWiseRoles []v1.ResourceWiseRoles
	for _, eachResource := range resourceWiseRoles {
		if _, ok := resourceMap[eachResource.Name]; ok {
			var addedRoles []v1.Role
			for _, eachRole := range eachResource.Roles {
				if val, roleOk := roleMap[eachRole.Name]; roleOk {
					addedRoles = append(addedRoles, val)
				}
			}
			if len(addedRoles) > 0 {
				resourceWiseRole := v1.ResourceWiseRoles{
					Name:  eachResource.Name,
					Roles: addedRoles,
				}
				newResourceWiseRoles = append(newResourceWiseRoles, resourceWiseRole)
			}
		}
	}
	return newResourceWiseRoles
}

// CheckDuplicateData checks and removes duplicate roles from v1.UserResourcePermission
func CheckDuplicateData(data v1.UserResourcePermission) v1.UserResourcePermission {
	resourceMap := make(map[string]int)
	temp := v1.UserResourcePermission{}
	for _, eachResource := range data.Resources {
		roleMap := make(map[string]int)
		if _, ok := resourceMap[eachResource.Name]; !ok {
			resourceMap[eachResource.Name] = len(temp.Resources)
			tempResource := v1.ResourceWiseRoles{Name: eachResource.Name}
			temp.Resources = append(temp.Resources, CheckDuplicateRoles(eachResource, tempResource, roleMap))
		} else {
			tempResource := v1.ResourceWiseRoles{Name: eachResource.Name}
			temp.Resources[resourceMap[eachResource.Name]] = CheckDuplicateRoles(eachResource, tempResource, roleMap)
		}
	}
	return temp
}

// CheckDuplicateRoles checks and removes duplicate roles from v1.UserResourcePermission and returns v1.ResourceWiseRoles
func CheckDuplicateRoles(resource v1.ResourceWiseRoles, tempResource v1.ResourceWiseRoles, roleMap map[string]int) v1.ResourceWiseRoles {
	for _, eachRole := range resource.Roles {
		if _, ok := roleMap[eachRole.Name]; !ok {
			roleMap[eachRole.Name] = len(tempResource.Roles)
			tempResource.Roles = append(tempResource.Roles, v1.Role{Name: eachRole.Name})
		} else {
			tempResource.Roles[roleMap[eachRole.Name]] = v1.Role{Name: eachRole.Name}
		}
	}
	return tempResource
}
