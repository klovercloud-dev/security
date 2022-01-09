package service

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	v1 "github.com/klovercloud-ci/core/v1"
)

type Jwt interface {
 GetRsaKeys() *v1.RsaKeys
 GenerateToken(userUUID string,duration int,data interface{}) (string,string, error)
 IsTokenValid(tokenString string) (bool, *jwt.Token)
 GetPrivateKey() *rsa.PrivateKey
 GetPublicKey() *rsa.PublicKey
}