package api

import (
	"official/models"
	"github.com/astaxie/beego"
)

type BannerApiController struct {
	BaseController
}

func (self *BannerApiController) BannerList(){
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	filters := make([]interface{}, 0)
	result, count := models.BannerGetListApi(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["img"] = "http://" + beego.AppConfig.String("dnsname") + v.Path
		row["link"] = v.Link
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
