package v1

import (
	"crypto/rsa"
	"github.com/klovercloud-ci/enums"
	"time"
)

type Resource struct {
	Name string `json:"name" bson:"name"`
}

type Permission struct {
	Name enums.PERMISION_TYPE `json:"name" bson:"name"`
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
	ID           string `json:"id" bson:"id"`
	FirstName    string `json:"first_name" bson:"first_name" `
	LastName     string `json:"last_name" bson:"last_name"`
	Email        string `json:"email" bson:"email" `
	Password     string `json:"password" bson:"password" `
	Status       string `json:"status" bson:"status"`
	CreatedDate  string `json:"created_date" bson:"created_date"`
	UpdatedDate  string `json:"updated_date" bson:"updated_date"`
	Token        string `json:"token" bson:"token"`
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
	AuthType     string `json:"auth_type" bson:"auth_type"`
}

type UserRegistrationDto struct {
	ID                 string                 `json:"_id" bson:"_id"`
	FirstName          string                 `json:"first_name" bson:"first_name" `
	LastName           string                 `json:"last_name" bson:"last_name"`
	Email              string                 `json:"email" bson:"email" `
	Password           string                 `json:"password" bson:"password" `
	Status             string                 `json:"status" bson:"status"`
	CreatedDate        string                 `json:"created_date" bson:"created_date"`
	UpdatedDate        string                 `json:"updated_date" bson:"updated_date"`
	Token              string                 `json:"token" bson:"token"`
	RefreshToken       string                 `json:"refresh_token" bson:"refresh_token"`
	AuthType           string                 `json:"auth_type" bson:"auth_type"`
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
		Password:     u.Password,
		Status:       u.Status,
		CreatedDate:  u.CreatedDate,
		UpdatedDate:  u.UpdatedDate,
		Token:        u.Token,
		RefreshToken: u.RefreshToken,
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
	Email          string           `json:"email" bson:"email"`
	Password        string           `json:"password" bson:"password"`
}

type JWTPayLoad struct {
	AccessToken  string           `json:"access_token" bson:"access_token"`
	RefreshToken string           `json:"refresh_token" bson:"refresh_token"`
	ExpiresIn int64    `json:"expires_in" bson:"expires_in"`
	CreationTime time.Time `json:"creation_time" bson:"creation_time"`
}