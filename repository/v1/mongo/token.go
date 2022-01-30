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

// TokenCollection collection name
var (
	TokenCollection = "tokenCollection"
)

type tokenRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (t tokenRepository) GetByUID(uid string) v1.Token {
	var res v1.Token
	query := bson.M{
		"$and": []bson.M{},
	}
	and := []bson.M{{"uid": uid}}
	query["$and"] = and
	coll := t.manager.Db.Collection(TokenCollection)
	result, err := coll.Find(t.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Token)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return v1.Token{}
		}
		res = *elemValue
	}
	return res
}

func (t tokenRepository) Store(token v1.Token) error {
	coll := t.manager.Db.Collection(TokenCollection)
	_, err := coll.InsertOne(t.manager.Ctx, token)
	if err != nil {
		log.Println("[ERROR] Insert document:", err.Error())
	}
	return nil
}

func (t tokenRepository) Delete(uid string) error {
	coll := t.manager.Db.Collection(TokenCollection)
	filter := bson.M{"uid": uid}
	res, err := coll.DeleteOne(t.manager.Ctx, filter)
	if err != nil {
		log.Println("[ERROR]", err)
	}
	if res.DeletedCount == 0 {
		return errors.New("[ERROR] Delete failed")
	}
	return err
}

func (t tokenRepository) Update(token string, refreshToken string, existingToken string) error {
	oldTokenObj := t.GetByToken(existingToken)
	if oldTokenObj.Uid == "" {
		return errors.New("[ERROR] Token does not exists")
	}
	oldTokenObj.Token = token
	oldTokenObj.RefreshToken = refreshToken

	filter := bson.M{
		"$and": []interface{}{
			bson.M{"uid": oldTokenObj.Uid},
		},
	}
	update := bson.M{
		"$set": oldTokenObj,
	}
	upsert := false
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := t.manager.Db.Collection(TokenCollection)
	err := coll.FindOneAndUpdate(t.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (t tokenRepository) GetByToken(token string) v1.Token {
	var res v1.Token
	query := bson.M{
		"$or": []interface{}{
			bson.M{"token": token},
			bson.M{"refresh_token": token},
		},
	}
	coll := t.manager.Db.Collection(TokenCollection)
	result, err := coll.Find(t.manager.Ctx, query, nil)
	if err != nil {
		log.Println(err.Error())
	}
	for result.Next(context.TODO()) {
		elemValue := new(v1.Token)
		err := result.Decode(elemValue)
		if err != nil {
			log.Println("[ERROR]", err)
			return v1.Token{}
		}
		res = *elemValue
	}
	return res
}

// NewTokenRepository returns repository.Token type repository
func NewTokenRepository(timeout int) repository.Token {
	return &tokenRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
