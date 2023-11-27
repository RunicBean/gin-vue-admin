// 自动生成模板Company
package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 公司 结构体  Company
type Company struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`  //公司名称 
      Path  string `json:"path" form:"path" gorm:"column:path;comment:;"`  //前端路径ID 
}


// TableName 公司 Company自定义表名 company
func (Company) TableName() string {
  return "company"
}

