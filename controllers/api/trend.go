package api

import (
	"official/models"
	"github.com/astaxie/beego"
)

type TrendApiController struct {
	BaseController
}

func (self *TrendApiController) Trend() {
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	result, count := models.NewsGetTrend(page, limit)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["subTitle"] = v.SubTitle
		row["createTime"] = v.CreateTime.String()[:19]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *TrendApiController) TrendCover() {
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	filters := make([]interface{}, 0)

	filters = append(filters, "status", 1)
	filters = append(filters, "genre", 1)
	filters = append(filters, "cover__icontains", ".")

	result, count := models.NewsGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))

	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["title"] = v.Title
		row["subTitle"] = v.SubTitle
		//row["cover"] = "http://" + libs.GetPulicIP()+":"+beego.AppConfig.String("httpport") + v.Cover
		row["cover"] = "http://" + beego.AppConfig.String("dnsname") + v.Cover
		row["createTime"] = v.CreateTime.String()[:19]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
