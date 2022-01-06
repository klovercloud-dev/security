package mongo

import (
	"context"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

// PermissionCollection collection name
var (
	PermissionCollection = "permissionCollection"
)

type permissionRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (p permissionRepository) Store(permission v1.Permission) error {
	coll := p.manager.Db.Collection(PermissionCollection)
	_, err := coll.InsertOne(p.manager.Ctx, permission)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (p permissionRepository) Get() []v1.Permission {
	var permissions []v1.Permission
	coll := p.manager.Db.Collection(PermissionCollection)
	result, err := coll.Find(p.manager.Ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		var permissionValues []v1.Permission
		err := result.Decode(permissionValues)
		if err != nil {
			log.Println("[ERROR]", err)
			return nil
		}
		for _, each := range permissionValues {
			permissions = append(permissions, each)
		}
	}
	return permissions
}

func (p permissionRepository) Delete(permissionName string) error {
	coll := p.manager.Db.Collection(PermissionCollection)
	filter := bson.M{"PermissionName": permissionName}
	_, err := coll.DeleteOne(p.manager.Ctx, filter)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	return err
}

func newPermissionRepository(manager *dmManager, timeout time.Duration) repository.Permission {
	return &permissionRepository{
		manager: manager,
		timeout: timeout,
	}
}
