// 自动生成模板ItemData
package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 项目数据 结构体  ItemData
type ItemData struct {
      global.GVA_MODEL
      ImageLink  string `json:"imageLink" form:"imageLink" gorm:"column:image_link;comment:;"`  //图片链接 
      ExtLink  string `json:"extLink" form:"extLink" gorm:"column:ext_link;comment:;"`  //外部链接 
      Company  string `json:"company" form:"company" gorm:"column:company;comment:;"`  //公司 
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;"`  //标题 
}


// TableName 项目数据 ItemData自定义表名 item_data
func (ItemData) TableName() string {
  return "item_data"
}

