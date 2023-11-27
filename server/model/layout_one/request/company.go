package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/layout_one"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CompanySearch struct{
    layout_one.Company
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
