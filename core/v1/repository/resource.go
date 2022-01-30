package repository

import v1 "github.com/klovercloud-ci/core/v1"

// Resource Repository operations resource.
type Resource interface {
	Store(resource v1.Resource) error
	Get() []v1.Resource
	GetByName(name string) v1.Resource
	Delete(resourceName string) error
}
