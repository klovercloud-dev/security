package api

import (
	"github.com/labstack/echo/v4"
)

//Oauth api operations
type Oauth interface {
	Login(context echo.Context) error
}
