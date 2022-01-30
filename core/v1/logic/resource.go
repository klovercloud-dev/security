package logic

import (
	"errors"
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type resourceService struct {
	repo repository.Resource
}

func (r resourceService) Get() []v1.Resource {
	return r.repo.Get()
}

func (r resourceService) Store(resource v1.Resource) error {
	resources := r.repo.Get()
	for _, res := range resources {
		if res.Name == resource.Name {
			return errors.New("Resource already exists!")
		}
	}
	return r.repo.Store(resource)
}

func (r resourceService) GetByName(name string) v1.Resource {
	return r.repo.GetByName(name)
}

func (r resourceService) Delete(name string) error {
	return r.repo.Delete(name)
}

// NewResourceService returns service.Resource type service
func NewResourceService(repo repository.Resource) service.Resource {
	return &resourceService{
		repo: repo,
	}
}
