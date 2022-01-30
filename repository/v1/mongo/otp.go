package mongo

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// OtpCollection collection name
var (
	OtpCollection = "otpCollection"
)

type otpRepository struct {
	manager *dmManager
	timeout time.Duration
}

func (o otpRepository) Store(otp v1.Otp) error {
	otp.Exp = time.Now().UTC().Add(time.Minute * 5)
	filter := bson.M{"id": otp.ID}
	update := bson.M{
		"$set": otp,
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	coll := o.manager.Db.Collection(OtpCollection)
	err := coll.FindOneAndUpdate(o.manager.Ctx, filter, update, &opt)
	if err != nil {
		log.Println("[ERROR]", err.Err())
	}
	return nil
}

func (o otpRepository) FindByOtp(otp string) v1.Otp {
	elemValue := new(v1.Otp)
	filter := bson.M{"otp": otp}
	coll := o.manager.Db.Collection(OtpCollection)
	result := coll.FindOne(o.manager.Ctx, filter)
	err := result.Decode(elemValue)
	if err != nil {
		log.Println("[ERROR]", err)
		return *elemValue
	}
	return *elemValue
}

// NewOtpRepository returns repository.Otp type repository
func NewOtpRepository(timeout int) repository.Otp {
	return &otpRepository{
		manager: GetDmManager(),
		timeout: time.Duration(timeout),
	}
}
