package api

import (
	"github.com/labstack/echo/v4"
)

type User interface {
	Registration(context echo.Context) error
	Get(context echo.Context) error
	GetByID(context echo.Context) error
	Delete(context echo.Context) error
    ResetPassword(context echo.Context) error
	ForgotPassword(context echo.Context) error
	AttachCompany(context echo.Context) error
	UpdateStatus(context echo.Context) error
	Update(context echo.Context) error
	UpdateUserResourcePermission(context echo.Context) error
}
