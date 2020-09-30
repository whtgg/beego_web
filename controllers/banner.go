package controllers

import (
	"official/models"
	"strings"
	"fmt"
)

type BannerController struct {
	BaseController
}

func (self *BannerController) List() {
	self.Data["pageTitle"] = "Banner列表"
	self.display()
}

func (self *BannerController) Add() {
	self.Data["pageTitle"] = "Banner添加"
	self.display()
}

func (self *BannerController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = self.pageSize
	}

	filters := make([]interface{},0)
	result, count := models.BannerGetList(page, limit,filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["path"] = v.Path
		row["status"] = v.Status
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *BannerController) Save() {
	Banner := new(models.Banner)
	Banner.Path = strings.TrimSpace(self.GetString("path"))
	Banner.Link = strings.TrimSpace(self.GetString("link"))
	Banner.Status, _ = self.GetInt8("status")
	keyId, _ := self.GetInt("id")
	if keyId == 0 {
		if _, err := models.BannerAdd(Banner); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		res, err := models.BannerGetById(keyId)
		if res != nil {
			res.Path = Banner.Path
			res.Status = Banner.Status
			res.Link = Banner.Link
			if err := res.BannerUpdate(); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BannerController) Edit() {
	self.Data["pageTitle"] = "Banner编辑"

	keyId, err := self.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}

	banner, _ := models.BannerGetById(keyId)
	list := make(map[string]interface{})
	list["id"] = banner.Id
	list["path"] = banner.Path
	list["link"] = banner.Link
	list["status"] = banner.Status
	self.Data["banner"] = list
	self.display()
}

func (self *BannerController) Delete() {
	keyId, _ := self.GetInt("id")
	res, _ := models.BannerGetById(keyId)
	res.Status = 9
	err := res.BannerUpdate()
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}
