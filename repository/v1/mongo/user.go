package mongo

import (
	"context"
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// UserCollection collection name
var (
	UserCollection = "userCollection"
)

type userRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (u userRepository) GetByEmail(email string) v1.User {
	var res v1.User
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"email": email}}
	query["$and"] = and
	coll := u.manager.Db.Collection(UserCollection)
	result, err := coll.Find(u.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
		return v1.User{}
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.User)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res
}

func (u userRepository) Store(user v1.User) error {
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	user.Password=string(hashedPassword)
	coll := u.manager.Db.Collection(UserCollection)
	_, err = coll.InsertOne(u.manager.Ctx, user)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (u userRepository) Get() []v1.User {
	var results []v1.User
	coll := u.manager.Db.Collection(UserCollection)
	result, err := coll.Find(u.manager.Ctx, bson.D{}, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.User)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		results = append(results, *elemValue)
	}
	return results
}

func (u userRepository) GetByID(id string) (v1.User, error) {
	var res v1.User
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"_id": id}}
	query["$and"] = and
	coll := u.manager.Db.Collection(UserCollection)
	result, err := coll.Find(u.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.User)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			break
		}
		res = *elemValue
	}
	return res, nil
}

func (u userRepository) Delete(id string) error {
	coll := u.manager.Db.Collection(UserCollection)
	filter := bson.M{"id": id}
	res, err := coll.DeleteOne(u.manager.Ctx, filter)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	if res.DeletedCount == 0 {
		return errors.New("[ERROR] Delete failed")
	}
	return err
}

func NewUserRepository(timeout int) repository.User {
	return &userRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
