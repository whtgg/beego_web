package api

import "official/models"

type ContractApiController struct {
	BaseController
}

func (self *ContractApiController) ContractList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	filters := make([]interface{}, 0)
	result, count := models.ContractGetListApi(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["content"] = v.Content
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}