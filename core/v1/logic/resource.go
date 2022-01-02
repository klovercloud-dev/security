package logic

import (
	v1 "github.com/klovercloud-ci/core/v1"
	"github.com/klovercloud-ci/core/v1/repository"
	"github.com/klovercloud-ci/core/v1/service"
)

type resourceService struct {
	repo repository.Resource
}

func (r resourceService) Get() ([]v1.Resource, int64) {
	//TODO implement me
	panic("implement me")
}

func (r resourceService) Store(resource v1.Resource) error {
	return r.repo.Store(resource)
}

func (r resourceService) GetByName(name string) (v1.Resource, error) {
	//TODO implement me
	panic("implement me")
}

func (r resourceService) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

// NewCompanyService returns Company type service
func NewResourceService(repo repository.Resource) service.Resource {
	return &resourceService{
		repo: repo,
	}
}
