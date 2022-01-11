package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION mongo as db
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// INMEMORY in memory storage as db
	DEVELOP  = ENVIRONMENT("DEVELOP")
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

var PERMISSION_LIST = []PERMISION_TYPE{
	PERMISSION_READ,
	PERMISSION_UPDATE,
	PERMISSION_DELETE,
	PERMISSION_CREATE,
}

type ROLE_UPDATE_OPTION string

const (
	APPEND_PERMISSION = ROLE_UPDATE_OPTION("append")
	REMOVE_PERMISSION = ROLE_UPDATE_OPTION("remove")
)

// TOKEN_TYPE token type of user
type TOKEN_TYPE string

const (
	// REGULAR_TOKEN refers to  limited lifetime token and refresh token
	REGULAR_TOKEN = TOKEN_TYPE("regular")
	// CTL_TOKEN refers to  long lifetime token and refresh token
	CTL_TOKEN     = TOKEN_TYPE("ctl")
)

// USER_UPDATE_ACTION users update action
type USER_UPDATE_ACTION string

const (
	// RESET_PASSWORD refers to password reset action
	RESET_PASSWORD = USER_UPDATE_ACTION("reset_password")
	// FORGOT_PASSWORD refers to password forgot action
	FORGOT_PASSWORD     = USER_UPDATE_ACTION("forgot_password")
)
