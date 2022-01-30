package v1

import (
	"crypto/rsa"
	"errors"
	"github.com/klovercloud-ci/enums"
	"net/mail"
	"time"
)

// Resource holds resource name.
type Resource struct {
	Name string `json:"name" bson:"name"`
}

// Permission holds permission names.
type Permission struct {
	Name string `json:"name" bson:"name"`
}

// Role holds role wise permissions.
type Role struct {
	Name        string       `json:"name" bson:"name"`
	Permissions []Permission `json:"permissions" bson:"permissions"`
}

// UserResourcePermissionDto holds metadata, user and resource wise permissions.
type UserResourcePermissionDto struct {
	Metadata  UserMetadata        `json:"metadata" bson:"-"`
	UserId    string              `json:"user_id" bson:"user_id"`
	Resources []ResourceWiseRoles `json:"resources" bson:"resources"`
}

// UserResourcePermission dto that holds metadata, user and resource wise roles.
type UserResourcePermission struct {
	Metadata  UserMetadata           `json:"metadata" bson:"-"`
	UserId    string                 `json:"user_id" bson:"user_id"`
	Resources []ResourceWiseRolesDto `json:"resources" bson:"resources"`
}

// ResourceWiseRoles holds resource wise roles.
type ResourceWiseRoles struct {
	Name  string `json:"name" bson:"name"`
	Roles []Role `json:"roles" bson:"roles"`
}

// ResourceWiseRolesDto dto that holds resource wise role dtos.
type ResourceWiseRolesDto struct {
	Name  string    `json:"name" bson:"name"`
	Roles []RoleDto `json:"roles" bson:"roles"`
}

// RoleDto dto that holds role name.
type RoleDto struct {
	Name string `json:"name" bson:"name"`
}

// User holds users info.
type User struct {
	Metadata           UserMetadata           `json:"metadata" bson:"metadata"`
	ID                 string                 `json:"id" bson:"id"`
	FirstName          string                 `json:"first_name" bson:"first_name" `
	LastName           string                 `json:"last_name" bson:"last_name"`
	Email              string                 `json:"email" bson:"email" `
	Phone              string                 `json:"phone" bson:"phone" `
	Password           string                 `json:"password" bson:"password" `
	Status             enums.STATUS           `json:"status" bson:"status"`
	CreatedDate        time.Time              `json:"created_date" bson:"created_date"`
	UpdatedDate        time.Time              `json:"updated_date" bson:"updated_date"`
	AuthType           enums.AUTH_TYPE        `json:"auth_type" bson:"auth_type"`
	ResourcePermission UserResourcePermission `json:"resource_permission" bson:"resource_permission"`
}

// UserMetadata holds users metadata.
type UserMetadata struct {
	CompanyId string `json:"company_id" bson:"company_id"`
}

// UserRegistrationDto dto that holds user registration info.
type UserRegistrationDto struct {
	Metadata           UserMetadata           `json:"metadata"`
	ID                 string                 `json:"id" bson:"id"`
	FirstName          string                 `json:"first_name" bson:"first_name" `
	LastName           string                 `json:"last_name" bson:"last_name"`
	Email              string                 `json:"email" bson:"email" `
	Phone              string                 `json:"phone" bson:"phone"`
	Password           string                 `json:"password" bson:"password" `
	Status             enums.STATUS           `json:"status" bson:"status"`
	CreatedDate        time.Time              `json:"created_date" bson:"created_date"`
	UpdatedDate        time.Time              `json:"updated_date" bson:"updated_date"`
	AuthType           enums.AUTH_TYPE        `json:"auth_type" bson:"auth_type"`
	ResourcePermission UserResourcePermission `json:"resource_permission" bson:"resource_permission"`
}

// RoleUpdateOption contains options for role update.
type RoleUpdateOption struct {
	Option enums.ROLE_UPDATE_OPTION `json:"option" bson:"option"`
}

// GetUserFromUserRegistrationDto converts User from UserRegistrationDto
func GetUserFromUserRegistrationDto(u UserRegistrationDto) User {
	user := User{
		Metadata:           UserMetadata{CompanyId: u.Metadata.CompanyId},
		ID:                 u.ID,
		FirstName:          u.FirstName,
		LastName:           u.LastName,
		Email:              u.Email,
		Phone:              u.Phone,
		Password:           u.Password,
		Status:             u.Status,
		CreatedDate:        u.CreatedDate,
		UpdatedDate:        u.UpdatedDate,
		AuthType:           u.AuthType,
		ResourcePermission: u.ResourcePermission,
	}
	return user
}

// RsaKeys contains RSA keys.
type RsaKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// Token contains token info.
type Token struct {
	Uid          string           `json:"uid" bson:"uid"`
	Token        string           `json:"token" bson:"token"`
	RefreshToken string           `json:"refresh_token" bson:"refresh_token"`
	Type         enums.TOKEN_TYPE `json:"type" bson:"type"`
}

// LoginDto contains user login info.
type LoginDto struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

// RefreshTokenDto contains refresh token.
type RefreshTokenDto struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

// JWTPayLoad contains payload of JWT token.
type JWTPayLoad struct {
	AccessToken  string `json:"access_token" bson:"access_token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}

// PasswordResetDto contains data for password reset
type PasswordResetDto struct {
	Otp             string `json:"otp" bson:"otp"`
	Email           string `json:"email" bson:"email"`
	CurrentPassword string `json:"current_password" bson:"current_password"`
	NewPassword     string `json:"new_password" bson:"new_password"`
}

// Validate validates UserRegistrationDto data
func (u UserRegistrationDto) Validate() error {
	if u.ID == "" {
		return errors.New("user id is required")
	}
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return err
	}
	if u.AuthType != enums.PASSWORD {
		return errors.New("invalid user AuthType")
	}
	return u.ResourcePermission.Validate()
}

// Validate validates UserResourcePermissionDto data
func (u UserResourcePermission) Validate() error {
	for _, eachResource := range u.Resources {
		if eachResource.Name == "" {
			return errors.New("[ERROR]: Blank resource name")
		}
		if eachResource.Name != string(enums.USER) && eachResource.Name != string(enums.PIPELINE) && eachResource.Name != string(enums.PROCESS) && eachResource.Name != string(enums.COMPANY) && eachResource.Name != string(enums.REPOSITORY) && eachResource.Name != string(enums.APPLICATION) {
			return errors.New("[ERROR]: Invalid resource name")
		}
		for _, eachRole := range eachResource.Roles {
			if err := eachRole.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// Validate validates Role data
func (r RoleDto) Validate() error {
	if r.Name == "" {
		return errors.New("[ERROR]: Blank role name")
	}
	return nil
}

// Validate validates Permission data
func (p Permission) Validate() error {
	if p.Name == "" {
		return errors.New("[ERROR]: Blank permission name")
	}
	if p.Name != string(enums.CREATE) && p.Name != string(enums.READ) && p.Name != string(enums.UPDATE) && p.Name != string(enums.DELETE) {
		return errors.New("[ERROR]: Invalid permission name")
	}
	return nil
}

// Otp contains otp data
type Otp struct {
	ID    string    `json:"id" bson:"id"`
	Email string    `json:"email" bson:"email"`
	Phone string    `json:"phone" bson:"phone"`
	Otp   string    `json:"otp" bson:"otp"`
	Exp   time.Time `json:"exp" bson:"exp"`
}

// Company contains company data
type Company struct {
	MetaData     CompanyMetadata `bson:"_metadata" json:"_metadata"`
	Id           string          `bson:"id" json:"id"`
	Name         string          `bson:"name" json:"name"`
	Repositories []Repository    `bson:"repositories" json:"repositories"`
}

// CompanyMetadata contains company metadata info
type CompanyMetadata struct {
	Labels                    map[string]string `bson:"labels" json:"labels" yaml:"labels"`
	NumberOfConcurrentProcess int64             `bson:"number_of_concurrent_process" json:"number_of_concurrent_process" yaml:"number_of_concurrent_process"`
	TotalProcessPerDay        int64             `bson:"total_process_per_day" json:"total_process_per_day" yaml:"total_process_per_day"`
}

// Repository contains repository info
type Repository struct {
	Id           string        `bson:"id" json:"id"`
	Type         string        `bson:"type" json:"type"`
	Token        string        `bson:"token" json:"token"`
	Applications []Application `bson:"applications" json:"applications"`
}

// Application contains application info
type Application struct {
	MetaData ApplicationMetadata `bson:"_metadata" json:"_metadata"`
	Url      string              `bson:"url" json:"url"`
	Webhook  GitWebhook          `bson:"webhook" json:"webhook"`
}

// ApplicationMetadata contains application metadata info
type ApplicationMetadata struct {
	Labels           map[string]string `bson:"labels" json:"labels"`
	Id               string            `bson:"id" json:"id"`
	Name             string            `bson:"name" json:"name"`
	IsWebhookEnabled bool              `bson:"is_webhook_enabled" json:"is_webhook_enabled"`
}

// GitWebhook contains github web hook data
type GitWebhook struct {
	Type   string   `json:"type"`
	ID     string   `json:"id"`
	Active bool     `json:"active"`
	Events []string `json:"events"`
	Config struct {
		URL         string `json:"url"`
		InsecureSsl string `json:"insecure_ssl"`
		ContentType string `json:"content_type"`
	} `json:"config"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	URL           string    `json:"url"`
	TestURL       string    `json:"test_url"`
	PingURL       string    `json:"ping_url"`
	DeliveriesURL string    `json:"deliveries_url"`
}
