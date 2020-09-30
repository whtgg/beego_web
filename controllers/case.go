/**
  * Created by: Jianyi
  * Date: 2019/1/4
  * Time: 17:36
  * Description:
 **/

package controllers

import (
	"official/models"
	"strings"
	"fmt"
)

type CasesController struct {
	BaseController
}

func (self *CasesController) List() {
	self.Data["pageTitle"] = "案例管理"
	self.display()
}

func (self *CasesController) Add() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "案例添加"
	self.Data["leaders"] = models.LeaderAll()
	self.display()
}

func (self *CasesController) Save() {
	Cases := new(models.Cases)
	Cases.Name = strings.TrimSpace(self.GetString("name"))
	Cases.CsrInfo = strings.TrimSpace(self.GetString("csr_info"))
	Cases.CaseInfo = strings.TrimSpace(self.GetString("case_info"))
	Cases.ProjectLeader,_ = self.GetInt("project_leader")
	Cases.Status,_ = self.GetInt("status")
	Cases.CaseType,_ = self.GetInt("case_type")

	keyId, _ := self.GetInt("id")
	if keyId == 0 {
		if _, err := models.CasesAdd(Cases); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		res, err := models.CasesGetById(keyId)
		if res != nil {
			res.Name = Cases.Name
			res.CsrInfo = Cases.CsrInfo
			res.CaseInfo = Cases.CaseInfo
			res.ProjectLeader = Cases.ProjectLeader
			res.Status = Cases.Status
			res.CaseType = Cases.CaseType
			if err := res.CasesUpdate(); err != nil {
				self.ajaxMsg(err.Error(), MSG_ERR)
			}
		} else {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *CasesController) Edit() {
	self.Data["Wangedit"] = true
	self.Data["pageTitle"] = "案例编辑"
	self.Data["leaders"] = models.LeaderAll()

	keyId, err := self.GetInt("id")
	if err != nil {
		fmt.Println(err)
	}

	cases, _ := models.CasesGetById(keyId)
	list := make(map[string]interface{})
	list["id"] = cases.Id
	list["name"] = cases.Name
	list["csrInfo"] = cases.CsrInfo
	list["caseInfo"] = cases.CaseInfo
	list["caseType"] = cases.CaseType
	list["projectLeader"] = cases.ProjectLeader
	list["status"] = cases.Status
	self.Data["cases"] = list
	self.display()
}

func (self *CasesController) Table() {
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
	filters = append(filters, "status__lt", 9)
	result, count := models.CasesGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["csrName"] = v.Name
		row["csrInfo"] = v.CsrInfo
		row["caseInfo"] = v.CaseInfo
		leader,_ := models.LeaderGetById(v.ProjectLeader)
		row["projectLeader"] = leader.Name
		row["status"] = v.Status
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *CasesController) Delete() {
	keyId, _ := self.GetInt("id")
	res, _ := models.CasesGetById(keyId)
	res.Status = 9;
	err := res.CasesUpdate()
	if err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}
