package logic

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/dgrijalva/jwt-go"
	"github.com/klovercloud-ci/config"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/service"
	"log"
	"time"
)

var RsaKeys *v1.RsaKeys = nil

type jwtService struct {

}

func (j jwtService) GetRsaKeys() *v1.RsaKeys {
	if RsaKeys == nil {
		RsaKeys = &v1.RsaKeys{
			PrivateKey: j.GetPrivateKey(),
			PublicKey: j.GetPublicKey(),
		}
	}
	return RsaKeys
}

func (j jwtService) GenerateToken(userUUID string, duration int64, data interface{}) (string, string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims = jwt.MapClaims{
		"exp": duration,
		"iat": time.Now().Unix(),
		"sub": userUUID,
		"data":data,
	}
	tokenString, err := token.SignedString(RsaKeys.PrivateKey)
	if err != nil {
		return "","",err
	}
	token.Claims = jwt.MapClaims{
		"exp": duration+duration/4,
		"iat": time.Now().Unix(),
		"sub": userUUID,
	}
	refreshTokenStr,err:=token.SignedString(RsaKeys.PrivateKey)
	if err != nil {
		return "","",err
	}

	return tokenString, refreshTokenStr,nil
}

func (j jwtService) IsTokenValid(tokenString string) (bool, *jwt.Token) {
	claims := jwt.MapClaims{}
	 jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return (RsaKeys.PublicKey), nil
	})
	return false,nil
}

func (j jwtService) GetPrivateKey() *rsa.PrivateKey {
	block, rest := pem.Decode([]byte(config.PrivateKey))
	if rest != nil {
		log.Print(rest)
	}
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return privateKeyImported
}

func (j jwtService) GetPublicKey() *rsa.PublicKey {
	block, rest := pem.Decode([]byte(config.Publickey))
	if rest != nil {
		log.Print(rest)
	}
	publicKeyImported, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Print(err.Error())
	}
	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)
	if !ok {
		log.Println(err.Error())
	}
	return rsaPub
}

func NewJwtService() service.Jwt {
	return &jwtService{
	}
}
