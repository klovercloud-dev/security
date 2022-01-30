package service

import v1 "github.com/klovercloud-ci/core/v1"

// Token business operations.
type Token interface {
	Store(token v1.Token) error
	Delete(uid string) error
	Update(token string, refreshToken string, existingToken string) error
	GetByToken(token string) v1.Token
}
