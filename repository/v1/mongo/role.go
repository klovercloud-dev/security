package mongo

import (
	"context"
	"errors"
	v1 "github.com/klovercloud-ci-cd/security/core/v1"
	"github.com/klovercloud-ci-cd/security/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// RoleCollection collection name
var (
	RoleCollection = "roleCollection"
)

type roleRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (r roleRepository) Update(name string, permissions []v1.Permission) error {
	role := r.GetByName(name)
	role.Permissions = permissions
	filter := bson.M{
		"$and": []bson.M{
			{"name": name},
		},
	}
	update := bson.M{
		"$set": role,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := r.manager.Db.Collection(RoleCollection)
	err := coll.FindOneAndUpdate(r.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (r roleRepository) Store(role v1.RoleDto) error {
	if r.GetByName(role.Name).Name == "" {
		coll := r.manager.Db.Collection(RoleCollection)
		_, err := coll.InsertOne(r.manager.Ctx, role)
		if err != nil {
			log.Println("[ERROR] Insert document:", err.Error())
		}
	}
	return nil
}

func (r roleRepository) Get() []v1.RoleDto {
	var roles []v1.RoleDto
	coll := r.manager.Db.Collection(RoleCollection)
	result, err := coll.Find(r.manager.Ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.RoleDto)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return nil
		}
		roles = append(roles, *elemValue)
	}
	return roles
}

func (r roleRepository) GetByName(name string) v1.RoleDto {
	elemValue := new(v1.RoleDto)
	filter := bson.M{"name": name}
	coll := r.manager.Db.Collection(RoleCollection)
	result := coll.FindOne(r.manager.Ctx, filter)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
		return *elemValue
	}
	return *elemValue
}

func (r roleRepository) Delete(roleName string) error {
	coll := r.manager.Db.Collection(RoleCollection)
	filter := bson.M{"name": roleName}
	data, err := coll.DeleteOne(r.manager.Ctx, filter)
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

func (r roleRepository) AppendPermissions(name string, permissions []v1.Permission) error {
	role := r.GetByName(name)
	for i := range permissions {
		flag := false
		for j := range role.Permissions {
			if permissions[i].Name == role.Permissions[j].Name {
				flag = true
				break
			}
		}
		if !flag {
			role.Permissions = append(role.Permissions, permissions[i])
		}
	}
	err := r.Update(name, role.Permissions)
	if err != nil {
		return err
	}
	return nil
}

func (r roleRepository) RemovePermissions(name string, permissions []v1.Permission) error {
	role := r.GetByName(name)
	var newPermissions []v1.Permission
	for i := range role.Permissions {
		flag := false
		for j := range permissions {
			if role.Permissions[i].Name == permissions[j].Name {
				flag = true
				break
			}
		}
		if !flag {
			newPermissions = append(newPermissions, role.Permissions[i])
		}
	}
	err := r.Update(name, newPermissions)
	if err != nil {
		return err
	}
	return nil
}

// NewRoleRepository returns repository.RoleDto type repository
func NewRoleRepository(timeout int) repository.Role {
	return &roleRepository{manager: GetDmManager(), timeout: time.Duration(timeout)}
}
