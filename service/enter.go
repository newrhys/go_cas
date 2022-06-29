package service

import (
	"wave-admin/service/cms"
	"wave-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	CmsServiceGroup cms.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)