package api

import (
	"github.com/labstack/echo/v4"
)


type Oauth interface {
	Login(context echo.Context) error
}




