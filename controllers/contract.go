package controllers

import (
	"fmt"
	"official/models"
	"strings"
)

type ContractController struct {
	BaseController
}

func (self *ContractController) List() {
	self.Data["pageTitle"] = "签约列表"
	self.display()
}

func (self *ContractController) Add() {
	self.Data["pageTitle"] = "签约添加"
	self.display()
}

func (self *ContractController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = self.pageSize
	}

	filters := make([]interface{},0)
	result, count := models.ContractGetList(page, limit,filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["content"] = v.Content
		row["status"] = v.Status
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *ContractController) Save() {
	Contract := new(models.Contract)
	Contract.Content = strings.TrimSpace(self.GetString("content"))
	Contract.Status, _ = self.GetInt8("status")
	keyId, _ := self.GetInt("id")
	if keyId == 0 {
		if _, err := models.ContractAdd(Contract); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		res, err := models.ContractGetById(keyId)
		if res != nil {
			res.Content = Contract.Content
			res.Status = Contract.Status
			if err := res.ContractUpdate(); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *ContractController) Edit() {
	self.Data["pageTitle"] = "签约编辑"

	keyId, err := self.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}

	Contract, _ := models.ContractGetById(keyId)
	list := make(map[string]interface{})
	list["id"] = Contract.Id
	list["content"] = Contract.Content
	list["status"] = Contract.Status
	self.Data["contract"] = list
	self.display()
}

func (self *ContractController) Delete() {
	keyId, _ := self.GetInt("id")
	res, _ := models.ContractGetById(keyId)
	res.Status = 9
	err := res.ContractUpdate()
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}