package mongo

import (
	"context"
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/enums"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u userRepository) GetUsersByCompanyId(companyId string) []v1.User {
	panic("implement me")
}

func (u userRepository) UpdateStatus(id string, status enums.STATUS) error {
	panic("implement me")
}

func (u userRepository) AttachCompany(id, companyId string) error {
	user := u.GetByID(id)

	user.Metadata.CompanyId = companyId
	filter := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
	update := bson.M{
		"$set": user,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := u.manager.Db.Collection(UserCollection)
	err := coll.FindOneAndUpdate(u.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Err())
	}
	return nil
}

func (u userRepository) GetByPhone(phone string) v1.User {
	var res v1.User
	query := bson.M{
		"$and": []bson.M{
			{"phone": phone},
		},
	}
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

func (u userRepository) GetByToken(token string) v1.User {
	var res v1.User
	query := bson.M{
		"$or": []interface{}{
			bson.M{"token": token},
			bson.M{"refresh_token": token},
		},
	}
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
			return v1.User{}
		}
		res = *elemValue
	}
	return res
}

func (u userRepository) UpdatePassword(user v1.User) error {
	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	user.Password=string(hashedPassword)
	filter := bson.M{
		"$and": []interface{}{
			bson.M{"id": user.ID},
		},
	}
	update := bson.M{
		"$set": user,
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := u.manager.Db.Collection(UserCollection)
	uopdateErr := coll.FindOneAndUpdate(u.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", uopdateErr.Err())
	}
	return nil
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

func (u userRepository) GetByID(id string) v1.User{
	var res v1.User
	query := bson.M{
		"$and": []bson.M{
			{"id": id},
		},
	}
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
	return res
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
