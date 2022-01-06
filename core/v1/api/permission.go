package api

import "github.com/labstack/echo/v4"

//Permission resource api operations
type Permission interface {
	Store(context echo.Context) error
	Get(context echo.Context) error
	Delete(context echo.Context) error
}
