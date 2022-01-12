package api

import (
	"github.com/labstack/echo/v4"
)

type UserResourcePermission interface {
	Store(context echo.Context) error
	Get(context echo.Context) error
	GetByUserID(context echo.Context) error
	Delete(context echo.Context) error
	Update(context echo.Context) error
}
