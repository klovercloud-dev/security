package repository

import v1 "github.com/klovercloud-ci/core/v1"

type Resource interface {
	Store(resource v1.Resource) error
	Get() ([]v1.Resource, error)
	GetByName(name string) (v1.Resource, error)
	Delete(name string) error
}
