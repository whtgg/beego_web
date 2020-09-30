/**
  * Created by: Jianyi
  * Date: 2019/1/4
  * Time: 17:34
  * Description:
 **/
package api

import (
	"official/models"
	"fmt"
	"github.com/astaxie/beego"
)

var CaseType = map[string]int{
	"A":1,
	"F":2,
	"G":3,
	"U":4,
	"E":5,
}

type CasesApiController struct {
	BaseController
}

func (self *CasesApiController) CasesList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	filters := make([]interface{}, 0)
	result, count := models.CasesGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["projectLeader"] = v.ProjectLeader
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *CasesApiController) LeaderList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = PAGE
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = PAGE_SIZE
	}
	filters := make([]interface{}, 0)
	result, count := models.LeaderGetList(page, limit, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["proField"] = v.ProField
		row["wechat"] = v.Wechat
		row["wechat_img"] = v.WechatImg
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (self *CasesApiController) CasesDetailByLeaderId() {
	id, err := self.GetInt(":id")
	if err != nil {
		fmt.Println(err)
		self.ajaxList("失败", MSG_ERR, 0, 0)
	}
	result, _ := models.CasesLeaderGetByLeaderId(id)
	count:=len(result)
	list := make([]map[string]interface{}, count)
	//fmt.Println(result)
	for k, v :=range result {
		fmt.Println(v)
		cases := make(map[string]interface{})
		cases["csrName"] = v.CName
		cases["csrInfo"] = v.CsrInfo
		cases["caseInfo"] = v.CaseInfo
		cases["projectLeader"] = v.LName
		cases["proField"] = v.ProField
		cases["wechat"] = v.Wechat
		cases["wechat_img"] = "http://" + beego.AppConfig.String("dnsname") +v.WechatImg
		list[k]=cases
	}
	self.ajaxList("成功", MSG_OK, int64(count), list)
}


func (self *CasesApiController) CasesDetailByType() {
	typeStr := self.GetString("caseKey")
	if len(typeStr) == 0 {
		self.ajaxList("失败", MSG_ERR, 0, 0)
	}
	result, _ := models.CasesInfoGetByType(CaseType[typeStr])
	count:=len(result)
	list := make([]map[string]interface{}, count)
	for k, v :=range result {
		cases := make(map[string]interface{})
		cases["csrName"] = v.CName
		cases["csrInfo"] = v.CsrInfo
		cases["caseInfo"] = v.CaseInfo
		cases["projectLeader"] = v.LName
		cases["proField"] = v.ProField
		cases["wechat"] = v.Wechat
		cases["wechat_img"] = "http://" + beego.AppConfig.String("dnsname") +v.WechatImg
		cases["avatar"] = "http://" + beego.AppConfig.String("dnsname") + v.Avatar
		list[k]=cases
	}
	self.ajaxList("成功", MSG_OK, int64(count), list)
}
