package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION mongo as db
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// INMEMORY in memory storage as db
	DEV  = ENVIRONMENT("DEV")
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)

type PERMISION_TYPE string

const (
	PERMISSION_READ   = PERMISION_TYPE("read")
	PERMISSION_UPDATE = PERMISION_TYPE("update")
	PERMISSION_DELETE = PERMISION_TYPE("delete")
	PERMISSION_CREATE = PERMISION_TYPE("create")
)

type ROLE_TYPE string

const (
	ROLE_ADMIN       = ROLE_TYPE("admin")
	ROLE_USER        = ROLE_TYPE("user")
	ROLE_SUPER_ADMIN = ROLE_TYPE("super_admin")
	ROLE_MODERATOR   = ROLE_TYPE("moderator")
)

type ROLE_UPDATE_OPTION string

const (
	ROLE_UPDATE_OPTION_ADD    = ROLE_UPDATE_OPTION("add")
	ROLE_UPDATE_OPTION_REMOVE = ROLE_UPDATE_OPTION("remove")
)
