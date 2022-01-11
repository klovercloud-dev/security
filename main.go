package main

import (
	"github.com/klovercloud-ci/api"
	"github.com/klovercloud-ci/config"
	"github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/dependency"
)

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	go initResources()
	go initPermissions()
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
	resourceService:=dependency.GetV1PermissionService()
	for _,each:=range config.Resources {
		resourceService.Store(v1.Permission{Name: each})
	}
}