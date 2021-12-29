package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	ResourceRouter(g.Group("/resource"))
}

func ResourceRouter(g *echo.Group) {
	resourceApi := NewResourceApi(dependency.GetV1ResourceService())
	g.POST("", resourceApi.Store)
}
