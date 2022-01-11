package v1

import (
	"github.com/klovercloud-ci/dependency"
	"github.com/labstack/echo/v4"
)

func Router(g *echo.Group) {
	ResourceRouter(g.Group("/resource"))
	PermissionRouter(g.Group("/permission"))
	UserRouter(g.Group("/user"))
	RoleRouter(g.Group("/role"))
	UserResourcePermissionRouter(g.Group("/userResourcePermission"))
}

func ResourceRouter(g *echo.Group) {
	resourceApi := NewResourceApi(dependency.GetV1ResourceService())
	g.POST("", resourceApi.Store)
	g.GET("", resourceApi.Get)
	g.GET("/:resourceName", resourceApi.GetByName)
	g.DELETE("/:resourceName", resourceApi.Delete)
}

func PermissionRouter(g *echo.Group) {
	permissionApi := NewPermissionApi(dependency.GetV1PermissionService())
	g.POST("", permissionApi.Store)
	g.GET("", permissionApi.Get)
	g.DELETE("", permissionApi.Delete)
}

func RoleRouter(g *echo.Group) {
	roleApi := NewRoleApi(dependency.GetV1RoleService())
	g.POST("", roleApi.Store)
	g.GET("", roleApi.Get)
	g.GET("/:roleName", roleApi.GetByName)
	g.DELETE("/:roleName", roleApi.Delete)
	g.POST("/:roleName", roleApi.Update)
}

func UserRouter(g *echo.Group) {
	userApi := NewUserApi(dependency.GetV1UserService())
	userResourcePermissionApi := NewUserResourcePermissionApi(dependency.GetV1UserResourcePermissionService())
	g.POST("", userApi.Store)
	g.GET("", userApi.Get)
	g.GET("/:id", userApi.GetByID)
	g.DELETE("/:id", userApi.Delete)
	g.POST("/:id/userResourcePermission", userResourcePermissionApi.Update)
}

func UserResourcePermissionRouter(g *echo.Group) {
	userResourcePermissionApi := NewUserResourcePermissionApi(dependency.GetV1UserResourcePermissionService())
	g.POST("", userResourcePermissionApi.Store)
	g.GET("", userResourcePermissionApi.Get)
	g.GET("/:id", userResourcePermissionApi.GetByUserID)
	g.DELETE("/:id", userResourcePermissionApi.Delete)
}
