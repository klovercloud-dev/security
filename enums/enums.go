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

// ROLE_UPDATE_OPTION role update option type
type ROLE_UPDATE_OPTION string
const (
	// APPEND_PERMISSION refers to append permission update option
	APPEND_PERMISSION = ROLE_UPDATE_OPTION("append")
	// REMOVE_PERMISSION refers to remove permission update option
	REMOVE_PERMISSION = ROLE_UPDATE_OPTION("remove")
)

// TOKEN_TYPE token type of user
type TOKEN_TYPE string
const (
	// REGULAR_TOKEN refers to limited lifetime token and refresh token
	REGULAR_TOKEN = TOKEN_TYPE("regular")
	// CTL_TOKEN refers to long lifetime token and refresh token
	CTL_TOKEN = TOKEN_TYPE("ctl")
)

// USER_UPDATE_ACTION users update action
type USER_UPDATE_ACTION string

const (
	// RESET_PASSWORD refers to password reset action
	RESET_PASSWORD = USER_UPDATE_ACTION("reset_password")
	// FORGOT_PASSWORD refers to password forgot action
	FORGOT_PASSWORD     = USER_UPDATE_ACTION("forgot_password")
	// ATTACH_COMPANY refers to company attachment action
	ATTACH_COMPANY    = USER_UPDATE_ACTION("attach_company")
)

// STATUS status update action
type STATUS string

const (
	// ACTIVE user status for active user
	ACTIVE = STATUS("active")
	// INACTIVE user status for inactive user
	INACTIVE = STATUS("inactive")
	// DELETED user status for deleted user
	DELETED = STATUS("deleted")
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
	PHONE = MEDIA("phone")
)

// USER_REGISTRATION_ACTION user registration action
type USER_REGISTRATION_ACTION string
const (
	// CREATE_USER refers to create user by admin
	CREATE_USER = USER_REGISTRATION_ACTION("create_user")
)

// ROLE role string
type ROLE string
const (
	// ADMIN refers to admin role
	ADMIN = ROLE("ADMIN")
)

// RESOURCE resource string
type RESOURCE string
const (
	// USER refers to user resource
	USER = RESOURCE("user")
	// PIPELINE refers to pipeline resource
	PIPELINE = RESOURCE("pipeline")
	// PROCESS refers to process resource
	PROCESS = RESOURCE("process")
	// COMPANY refers to company resource
	COMPANY = RESOURCE("company")
	// REPOSITORY refers to repository resource
	REPOSITORY = RESOURCE("repository")
	// APPLICATION refers to application resource
	APPLICATION = RESOURCE("application")
)

// PERMISSION permission string
type PERMISSION string
const (
	// CREATE refers to CREATE permission
	CREATE = PERMISSION("CREATE")
	// READ refers to READ permission
	READ = PERMISSION("READ")
	// UPDATE refers to UPDATE permission
	UPDATE = PERMISSION("UPDATE")
	// DELETE refers to DELETE permission
	DELETE = PERMISSION("DELETE")
)