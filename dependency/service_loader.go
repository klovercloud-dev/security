package dependency

import (
	"github.com/klovercloud-ci/core/v1/logic"
	"github.com/klovercloud-ci/core/v1/service"
	"github.com/klovercloud-ci/repository/v1/mongo"
)

func GetV1ResourceService() service.Resource {
	var resource service.Resource
	resource = logic.NewResourceService(mongo.NewResourceRepository(300))

	return resource
}
