package main

import (
	"github.com/klovercloud-ci/api"
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/dependency"
	"github.com/klovercloud-ci/enums"
)

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	go initResources()
	initPermissions()
	initRoles()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}

func initResources(){
	resourceService:=dependency.GetV1ResourceService()
	for _,each:=range config.Resources {
		resourceService.Store(v1.Resource{Name: each})
	}
}


func initPermissions(){
	permissionService:=dependency.GetV1PermissionService()
	for _,each:=range config.Permissions {
		permissionService.Store(v1.Permission{Name: each})
	}
}

func initRoles(){
	permissions:=dependency.GetV1PermissionService().Get()
	role:=v1.Role{
		Name:      string(enums.ADMIN),
		Permissions: permissions,
	}
	dependency.GetV1RoleService().Store(role)
}