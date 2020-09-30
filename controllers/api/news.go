package api

import (
	"official/models"
	"fmt"
	"strings"
)

type NewsApiController struct {
	BaseController
}

/**
 * 新闻列表
 */
func (self *NewsApiController) NewsList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	filters := make([]interface{}, 0)
	tags := strings.TrimSpace(self.GetString("tags"))

	switch tags{
	case "new"://最新
		filters = append(filters,"tags",2)
	case "cm"://推荐
		filters = append(filters,"tags",3)
	case "hot"://热门
		filters = append(filters,"tags",4)
	default:
		filters = append(filters,"tags",1)
	}
	filters = append(filters, "status", 1)

	result, count := models.NewsGetList(page, limit, filters...)
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

/**
 *新闻详情
 */
func (self *NewsApiController) NewsDetail() {
	id, err := self.GetInt(":id")
	if err != nil {
		fmt.Println(err)
		self.ajaxList("失败", MSG_ERR, 0, 0)
	}
	result, _ := models.NewsGetById(id)
	news := make(map[string]interface{})
	news["content"] = result.Content
	news["title"] = result.Title
	news["subTitle"] = result.SubTitle
	news["createTime"] = result.CreateTime.String()[:19]
	prev,next,_ := models.NewsGetByIdRange(id)
	if prev.Id > 0 {
		prevs := make(map[string]interface{})
		prevs["id"] = prev.Id
		prevs["title"] = prev.Title
		news["prevItem"] = prevs
	}
	nexts := make(map[string]interface{})
	if next.Id > 0{
		nexts["id"] = next.Id
		nexts["title"] = next.Title
		news["nextItem"] = nexts
	}
	self.ajaxList("成功", MSG_OK, 0, news)
}

