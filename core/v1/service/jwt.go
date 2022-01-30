package service

import (
	"crypto/rsa"
	v1 "github.com/klovercloud-ci/core/v1"
)

// Jwt business operations.
type Jwt interface {
	GetRsaKeys() *v1.RsaKeys
	GenerateToken(userUUID string, duration int64, data interface{}) (token string, refreshToken string, err error)
	IsTokenValid(tokenString string) bool
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
}
