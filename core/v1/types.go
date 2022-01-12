package v1

import (
	"crypto/rsa"
	"errors"
	"github.com/klovercloud-ci/enums"
	"net/mail"
	"time"
)

type Resource struct {
	Name string `json:"name" bson:"name"`
}

type Permission struct {
	Name string `json:"name" bson:"name"`
}

type Role struct {
	Name        string       `json:"name" bson:"name"`
	Permissions []Permission `json:"permissions" bson:"permissions"`
}

type UserResourcePermission struct {
	UserId    string `json:"user_id" bson:"user_id"`
	Resources []struct {
		Name  string `json:"name" bson:"name"`
		Roles []Role `json:"roles" bson:"roles"`
	} `json:"resources" bson:"resources"`
}

type User struct {
	Metadata UserMetadata  `json:"metadata" bson:"metadata"`
	ID           string `json:"id" bson:"id"`
	FirstName    string `json:"first_name" bson:"first_name" `
	LastName     string `json:"last_name" bson:"last_name"`
	Email        string `json:"email" bson:"email" `
	Phone        string `json:"phone" bson:"phone" `
	Password     string `json:"password" bson:"password" `
	Status       enums.STATUS 		`json:"status" bson:"status"`
	CreatedDate  time.Time `json:"created_date" bson:"created_date"`
	UpdatedDate  time.Time `json:"updated_date" bson:"updated_date"`
	AuthType     enums.AUTH_TYPE 	`json:"auth_type" bson:"auth_type"`
}

type UserMetadata struct {
	CompanyId           string `json:"company_id" bson:"company_id"`
}
type UserRegistrationDto struct {
	ID                 string                 `json:"id" bson:"id"`
	FirstName          string                 `json:"first_name" bson:"first_name" `
	LastName           string                 `json:"last_name" bson:"last_name"`
	Email              string                 `json:"email" bson:"email" `
	Phone        	   string 				  `json:"phone" bson:"phone"`
	Password           string                 `json:"password" bson:"password" `
	Status             enums.STATUS           `json:"status" bson:"status"`
	CreatedDate        time.Time              `json:"created_date" bson:"created_date"`
	UpdatedDate        time.Time              `json:"updated_date" bson:"updated_date"`
	AuthType           enums.AUTH_TYPE        `json:"auth_type" bson:"auth_type"`
	ResourcePermission UserResourcePermission `json:"resource_permission" bson:"resource_permission"`
}

type RoleUpdateOption struct {
	Option enums.ROLE_UPDATE_OPTION `json:"option" bson:"option"`
}

func GetUserAndResourcePermissionBody(u UserRegistrationDto) (User, UserResourcePermission) {
	user := User{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Phone: 		  u.Phone,
		Password:     u.Password,
		Status:       u.Status,
		CreatedDate:  u.CreatedDate,
		UpdatedDate:  u.UpdatedDate,
		AuthType:     u.AuthType,
	}
	userResourcePermission := UserResourcePermission{
		UserId:    u.ID,
		Resources: u.ResourcePermission.Resources,
	}
	return user, userResourcePermission
}

type RsaKeys struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}
type Token struct {
	Uid          string           `json:"uid" bson:"uid"`
	Token        string           `json:"token" bson:"token"`
	RefreshToken string           `json:"refresh_token" bson:"refresh_token"`
	Type         enums.TOKEN_TYPE `json:"type" bson:"type"`
}

type LoginDto struct {
	Email          string     `json:"email" bson:"email"`
	Password       string     `json:"password" bson:"password"`
}

type RefreshTokenDto struct {
	RefreshToken 	string   `json:"refresh_token" bson:"refresh_token"`
}


type JWTPayLoad struct {
	AccessToken  	string   `json:"access_token" bson:"access_token"`
	RefreshToken 	string   `json:"refresh_token" bson:"refresh_token"`
}


type PasswordResetDto struct {
	Otp              string                 `json:"otp" bson:"otp"`
	Email              string                 `json:"email" bson:"email"`
	CurrentPassword  string           `json:"current_password" bson:"current_password"`
	NewPassword string           `json:"new_password" bson:"new_password"`
}



// Validate validates UserRegistrationDto data
func(u UserRegistrationDto) Validate() error{
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
	if u.Password == "" {
		return errors.New("password is required")
	} else if len(u.Password) < 8 {
		return errors.New("password length must be at least 8")
	}
	if u.AuthType != enums.PASSWORD {
		return errors.New("invalid user AuthType")
	}
	return u.ResourcePermission.Validate()
}

// Validate validates UserResourcePermission data
func(u UserResourcePermission) Validate() error {
	for _, eachResource := range u.Resources {
		if eachResource.Name == "" {
			return errors.New("resource name is required")
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
func(r Role) Validate() error {
	if r.Name == "" {
		return errors.New("role name is required")
	}
	for _, each := range r.Permissions {
		if err := each.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates Permission data
func(p Permission) Validate() error {
	if p.Name == "" {
		return errors.New("permission name is required")
	}
	return nil
}

type Otp struct {
	ID           string 	`json:"id" bson:"id"`
	Email        string     `json:"email" bson:"email"`
	Phone        string     `json:"phone" bson:"phone"`
	Otp 		 string  	`json:"otp" bson:"otp"`
	Exp 		 time.Time  `json:"exp" bson:"exp"`
}