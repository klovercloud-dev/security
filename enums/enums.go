package enums

// ENVIRONMENT run environment
type ENVIRONMENT string

const (
	// PRODUCTION production environment
	PRODUCTION = ENVIRONMENT("PRODUCTION")
	// DEVELOP development environment
	DEVELOP  = ENVIRONMENT("DEVELOP")
	// TEST test environment
	TEST = ENVIRONMENT("TEST")
)

const (
	// MONGO mongo as db
	MONGO = "MONGO"
	// INMEMORY in memory storage as db
	INMEMORY = "INMEMORY"
)


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
	// UPDATE_USER_RESOURCE_PERMISSION refers to update update_user_resource_permission action
	UPDATE_USER_RESOURCE_PERMISSION    = USER_UPDATE_ACTION("update_user_resource_permission")
)

// STATUS status update action
type STATUS string

const (
	// ACTIVE user status for active user
	ACTIVE = STATUS("active")
	// INACTIVE user status for inactive user
	INACTIVE = STATUS("inactive")
)

// AUTH_TYPE AuthType update action
type AUTH_TYPE string

const (
	// PASSWORD grand_type of users authentication
	PASSWORD = AUTH_TYPE("password")
)
// MEDIA otp media
type MEDIA string
const (
	// EMAIL refers to email media
	EMAIL = MEDIA("email")
	// PHONE refers to phone media
	PHONE     = MEDIA("phone")
)
