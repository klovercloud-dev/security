package main

import "github.com/klovercloud-ci/config"

// @title integration-manager API
// @description integration-manager API
func main() {
	e := config.New()
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
