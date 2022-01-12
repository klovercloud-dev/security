package mongo

import (
	"context"
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// UserResourcePermissionCollection name
var (
	UserResourcePermissionCollection = "userResourcePermissionCollection"
)

type userResourcePermissionRepository struct {
	manager *dmManager
	timeout time.Duration
}

// Store stores given userResourcePermission to database
func (u userResourcePermissionRepository) Store(userResourcePermission v1.UserResourcePermission) error {
	coll := u.manager.Db.Collection(UserResourcePermissionCollection)
	_, err := coll.InsertOne(u.manager.Ctx, userResourcePermission)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

// Get gets all the userResourcePermission
func (u userResourcePermissionRepository) Get() []v1.UserResourcePermission {
	var userResourcePermissions []v1.UserResourcePermission
	coll := u.manager.Db.Collection(UserResourcePermissionCollection)
	result, err := coll.Find(u.manager.Ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.UserResourcePermission)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return nil
		}
		userResourcePermissions = append(userResourcePermissions, *elemValue)
	}
	return userResourcePermissions
}

// GetByUserID gets a userResourcePermission corresponding to the given userID
func (u userResourcePermissionRepository) GetByUserID(userID string) v1.UserResourcePermission {
	elemValue := new(v1.UserResourcePermission)
	filter := bson.M{"user_id": userID}
	coll := u.manager.Db.Collection(UserResourcePermissionCollection)
	result := coll.FindOne(u.manager.Ctx, filter)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
		return *elemValue
	}
	return *elemValue
}


// Delete deletes a userResourcePermission corresponding to the given userID
func (u userResourcePermissionRepository) Delete(userID string) error {
	coll := u.manager.Db.Collection(UserResourcePermissionCollection)
	filter := bson.M{"user_id": userID}
	data, err := coll.DeleteOne(u.manager.Ctx, filter)
	if err != nil {
		log.Println("[ERROR]", err)
		return err
	}
	if data.DeletedCount == 0 {
		log.Println("No data found to delete!")
		return errors.New("no data found to delete")
	}
	return err
}

// Update updates existing userResourcePermission with the given userResourcePermission
func (u userResourcePermissionRepository) Update(userResourcePermission v1.UserResourcePermission) error {
	filter := bson.M{"user_id": userResourcePermission.UserId}
	update := bson.M{
		"$set": userResourcePermission,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := u.manager.Db.Collection(UserResourcePermissionCollection)
	err := coll.FindOneAndUpdate(u.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func NewUserResourcePermissionRepository(timeout int) repository.UserResourcePermission {
	return &userResourcePermissionRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
