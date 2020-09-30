/**
  * Created by: Jianyi
  * Date: 2019/1/4
  * Time: 14:02
  * Description:
 **/
package controllers

import (
	"official/models"
	"strings"
	"fmt"
)

type LeaderController struct {
	BaseController
}

func (self *LeaderController) List() {
	self.Data["pageTitle"] = "负责人管理"
	self.display()
}

func (self *LeaderController) Add() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "负责人添加"
	self.display()
}

func (self *LeaderController) Save() {
	Leader := new(models.Leader)
	Leader.Name = strings.TrimSpace(self.GetString("name"))
	Leader.ProField = strings.TrimSpace(self.GetString("pro_field"))
	Leader.Wechat = strings.TrimSpace(self.GetString("wechat"))
	Leader.Avatar = strings.TrimSpace(self.GetString("avatar"))
	Leader.WechatImg = strings.TrimSpace(self.GetString("wechat_img"))
	Leader.Status,_ = self.GetInt("status")
	keyId, _ := self.GetInt("id")
	if keyId == 0 {
		if _, err := models.LeaderAdd(Leader); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		res, err := models.LeaderGetById(keyId)
		if res != nil {
			res.Name = Leader.Name
			res.ProField = Leader.ProField
			res.Wechat = Leader.Wechat
			res.Status = Leader.Status
			res.Avatar = Leader.Avatar
			res.WechatImg = Leader.WechatImg
			if err := res.LeaderUpdate(); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *LeaderController) Edit() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "负责人编辑"

	keyId, err := self.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}

	leaders, _ := models.LeaderGetById(keyId)
	list := make(map[string]interface{})
	list["id"] = leaders.Id
	list["name"] = leaders.Name
	list["proField"] = leaders.ProField
	list["wechat"] = leaders.Wechat
	list["status"] = leaders.Status
	list["avatar"] = leaders.Avatar
	list["wechat_img"] = leaders.WechatImg
	self.Data["leader"] = list
	self.display()
}

func (self *LeaderController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = self.pageSize
	}

	name := strings.TrimSpace(self.GetString("name"))

	filters := make([]interface{}, 0)
	if name != "" {
		filters = append(filters, "name__contains", name)
	}
	filters = append(filters, "status__lt", 5)
	result, count := models.LeaderGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["proField"] = v.ProField
		row["wechat"] = v.Wechat
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *LeaderController) Delete() {
	keyId, _ := self.GetInt("id")
	res, _ := models.LeaderGetById(keyId)
	res.Status = 5
	err := res.LeaderUpdate()
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}
