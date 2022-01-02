package repository

import v1 "github.com/klovercloud-ci/core/v1"

type User interface {
	Store(user v1.User) error
	Get() ([]v1.User, error)
	GetByID(id string) (v1.User, error)
	Delete(id string) error
}
