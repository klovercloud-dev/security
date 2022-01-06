package api

import (
	"github.com/labstack/echo/v4"
)

type User interface {
	Store(context echo.Context) error
	Get(context echo.Context) error
	GetByID(context echo.Context) error
	Delete(context echo.Context) error
}
