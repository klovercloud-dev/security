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

// MEDIA otp media
type MEDIA string
const (
	// EMAIL refers to email media
	EMAIL = MEDIA("email")
	// PHOME refers to phone media
	PHONE     = MEDIA("phone")
)
