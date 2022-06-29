package v1

import (
	"wave-admin/controller/v1/cms"
	"wave-admin/controller/v1/system"
)

type ApiGroup struct {
	SystemApiGroup 	system.ApiGroup
	CmsApiGroup 	cms.ApiGroup
}

var ApiGroupApp = new(ApiGroup)