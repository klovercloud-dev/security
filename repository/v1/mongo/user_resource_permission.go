package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"log"
	"time"
)

// UserResourcePermission collection name
var (
	UserResourcePermission = "userResourcePermissionCollection"
)

type userResourcePermissionRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (u userResourcePermissionRepository) Store(userResourcePermission v1.UserResourcePermission) error {
	coll := u.manager.Db.Collection(UserResourcePermission)
	_, err := coll.InsertOne(u.manager.Ctx, userResourcePermission)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (u userResourcePermissionRepository) Get() ([]v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) GetByUserID(userID string) (v1.UserResourcePermission, error) {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) Delete(userID string) error {
	//TODO implement me
	panic("implement me")
}

func (u userResourcePermissionRepository) Update(userResourcePermission v1.UserResourcePermission) error {
	//TODO implement me
	panic("implement me")
}

func NewUserResourcePermissionRepository(timeout int) repository.UserResourcePermission {
	return &userResourcePermissionRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
