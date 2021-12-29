package repository

import v1 "github.com/klovercloud-ci/core/v1"

type ResourceRepository interface {
	Store(resource v1.Resource) error
}
