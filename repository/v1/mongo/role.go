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

// RoleCollection collection name
var (
	RoleCollection = "roleCollection"
)

type roleRepository struct {
	manager *dmManager
	timeout time.Duration
}

// Update updates existing roles with given permissions
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

// Store stores given role to database
func (r roleRepository) Store(role v1.Role) error {
	coll := r.manager.Db.Collection(RoleCollection)
	_, err := coll.InsertOne(r.manager.Ctx, role)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

// Get gets all the roles
func (r roleRepository) Get() []v1.Role {
	var roles []v1.Role
	coll := r.manager.Db.Collection(RoleCollection)
	result, err := coll.Find(r.manager.Ctx, bson.M{})
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Role)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return nil
		}
		roles = append(roles, *elemValue)
	}
	return roles
}

// GetByName gets a role corresponding to the given name
func (r roleRepository) GetByName(name string) v1.Role {
	elemValue := new(v1.Role)
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

// Delete deletes a role corresponding to the given name
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

// AppendPermissions add permissions to the existing role permissions
func (r roleRepository) AppendPermissions(name string, permissions []v1.Permission) error {
	role := r.GetByName(name)
	for i, _ := range permissions {
		flag := false
		for j, _ := range role.Permissions {
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

// RemovePermissions removes the given permission from an existing role
func (r roleRepository) RemovePermissions(name string, permissions []v1.Permission) error {
	role := r.GetByName(name)
	var newPermissions []v1.Permission
	for i, _ := range role.Permissions {
		flag := false
		for j, _ := range permissions {
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

func NewRoleRepository(timeout int) repository.Role {
	return &roleRepository{manager: GetDmManager(), timeout: time.Duration(timeout)}
}