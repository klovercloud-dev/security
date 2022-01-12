package api

import (
	"github.com/labstack/echo/v4"
)

//Role role api operations
type Role interface {
	Store(context echo.Context) error
	Get(context echo.Context) error
	GetByName(context echo.Context) error
	Delete(context echo.Context) error
	Update(context echo.Context) error
}