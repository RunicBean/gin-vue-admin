package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/layout_one"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	Layout_oneApiGroup layout_one.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
