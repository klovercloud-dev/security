package config

import (
	"github.com/joho/godotenv"
	"github.com/klovercloud-ci-cd/security/enums"
	"log"
	"os"
	"strings"
)

// ServerPort refers to server port.
var ServerPort string

// DbServer refers to database server ip.
var DbServer string

// DbPort refers to database server port.
var DbPort string

// DbUsername refers to database name.
var DbUsername string

// DbPassword refers to database password.
var DbPassword string

// DatabaseConnectionString refers to database connection string.
var DatabaseConnectionString string

// DatabaseName refers to database name.
var DatabaseName string

// Database refers to database options.
var Database string

// PrivateKey refers to rsa private key .
var PrivateKey string

// Publickey refers to rsa public key .
var Publickey string

// RegularTokenLifetime refers to token lifetime of regular.
var RegularTokenLifetime string

// CTLTokenLifetime refers to token lifetime of ctl.
var CTLTokenLifetime string

// RunMode refers to run mode.
var RunMode string

// MailServerHostEmail refers to mail servers host email address.
var MailServerHostEmail string

// MailServerHostEmailSecret refers to mail servers host email address secret.
var MailServerHostEmailSecret string

// SmtpHost refers to mail servers smtp host addresss.
var SmtpHost string

// SmtpPort refers to mail servers smtp port.
var SmtpPort string

// Resources refers to resources.
var Resources []string

// Permissions refers to permission.
var Permissions []string

// ApiServerUrl refers to api servers base url.
var ApiServerUrl string


// InitEnvironmentVariables initializes environment variables
func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = string(enums.DEVELOP)
	}

	if RunMode != string(enums.PRODUCTION) {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}
	log.Println("RUN MODE:", RunMode)

	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database = os.Getenv("DATABASE")
	if Database == enums.MONGO {
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
	PrivateKey = os.Getenv("PRIVATE_KEY")
	Publickey = os.Getenv("PUBLIC_KEY")
	RegularTokenLifetime = os.Getenv("REGULAR_TOKEN_LIFETIME")
	CTLTokenLifetime = os.Getenv("CTL_TOKEN_LIFETIME")

	MailServerHostEmail = os.Getenv("MAIL_SERVER_HOST_EMAIL")
	MailServerHostEmailSecret = os.Getenv("MAIL_SERVER_HOST_EMAIL_SECRET")
	SmtpHost = os.Getenv("SMTP_HOST")
	SmtpPort = os.Getenv("SMTP_PORT")
	Resources = strings.Split(os.Getenv("RESOURCES"), ",")
	Permissions = strings.Split(os.Getenv("PERMISSIONS"), ",")
	ApiServerUrl = os.Getenv("API_SERVER_URL")
}
