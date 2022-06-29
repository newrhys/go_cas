package router

import (
	"wave-admin/router/cms"
	"wave-admin/router/system"
)

type RouterGroup struct {
	System 	system.RouterGroup
	Cms 	cms.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
