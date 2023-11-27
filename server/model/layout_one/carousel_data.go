// 自动生成模板CarouselData
package layout_one

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 轮播图数据 结构体  CarouselData
type CarouselData struct {
      global.GVA_MODEL
      ImageLink  string `json:"imageLink" form:"imageLink" gorm:"column:image_link;comment:;"`  //图片链接 
      Company  string `json:"company" form:"company" gorm:"column:company;comment:;"`  //公司 
}


// TableName 轮播图数据 CarouselData自定义表名 carousel_data
func (CarouselData) TableName() string {
  return "carousel_data"
}

