package v1

import (
	"github.com/klovercloud-ci/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Resource struct {
	Name string `json:"name" bson:"name"`
}

type Permission struct {
	PermissionName enums.PERMISION_TYPE `json:"permission_type" bson:"permission_type"`
}

type Role struct {
	Name        string       `json:"name" bson:"name"`
	Permissions []Permission `json:"permissions" bson:"permissions"`
}

type UserResourcePermission struct {
	UserId    string `json:"user_id" bson:"user_id"`
	Resources []struct {
		Name  string
		Roles []Role
	} `json:"resources" bson:"resources"`
}

type User struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName   string             `json:"first_name" bson:"first_name" binding:"required"`
	LastName    string             `json:"last_name" bson:"last_name" binding:"required"`
	Email       string             `json:"email" bson:"email" binding:"required"`
	Password    string             `json:"password" bson:"password" binding:"required"`
	Role        enums.ROLE_TYPE    `json:"role" bson:"role"`
	Status      string             `json:"status" bson:"status"`
	CreatedDate string             `json:"created_date" bson:"created_date"`
	UpdatedDate string             `json:"updated_date" bson:"updated_date"`
}

type RoleUpdateOption struct {
	Option enums.ROLE_UPDATE_OPTION `json:"option" bson:"option"`
}
