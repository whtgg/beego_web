/**
  * Created by: Jianyi
  * Date: 2019/1/4
  * Time: 17:36
  * Description:
 **/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Cases struct {
	Id            int       `pk:"auto"`
	Name          string    `orm:"size(50)"`
	CsrInfo       string    `orm:"size(100)"`
	CaseInfo      string
	CaseType      int
	Status        int
	ProjectLeader int       `orm:"ref(fk)"`
}

type CasesLeader struct {
	CName         string    `orm:"size(50)"`
	CsrInfo       string    `orm:"size(100)"`
	CaseInfo      string
	LName         string    `orm:"size(40)"`
	ProField      string    `orm:"size(200)"`
	Wechat        string    `orm:"size(20)"`
	WechatImg     string    `orm:"size(20)"`
	Avatar		  string    `orm:"size(255)"`
}


func (this *Cases) TableName() string {
	return TableName("case")
}

func CasesAdd(record *Cases) (int64, error) {
	return orm.NewOrm().Insert(record)
}

func (this *Cases) CasesUpdate(fields ... string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func CasesGetList(page, pageSize int, filters ...interface{}) ([]*Cases, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Cases, 0)
	query := orm.NewOrm().QueryTable(TableName("case"))
	if len(filters) > 0 && filters != nil {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			fmt.Println(filters[k].(string), filters[k+1])
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("project_leader", "id").Limit(pageSize, offset).All(&list)
	return list, total
}


func CasesGetById(id int) (*Cases, error) {
	cases := new(Cases)
	err := orm.NewOrm().QueryTable(TableName("case")).Filter("id", id).One(cases)
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func CasesLeaderGetByLeaderId(id int) ([]*CasesLeader, error) {
	var lists []*CasesLeader
	_,err := orm.NewOrm().Raw(fmt.Sprintf("select c.name c_name,csr_info,case_info,l.name l_name,pro_field,wechat,wechat_img,avatar from pp_case c inner join pp_leader l on project_leader=l.id where  l.id=%d order by l.name asc, c.name asc", id)).QueryRows(&lists)

	if err != nil {
		return nil, err
	}

	return lists, nil
}


func CasesInfoGetByType(caseType int) ([]*CasesLeader, error) {
	var lists []*CasesLeader
	_,err := orm.NewOrm().Raw(fmt.Sprintf("select c.name c_name,csr_info,case_info,l.name l_name,pro_field,wechat ,avatar,wechat_img from pp_case c inner join pp_leader l on project_leader=l.id where  c.case_type=%d and c.status != 9 order by l.name asc, c.name asc",caseType)).QueryRows(&lists)
	if err != nil {
		return nil, err
	}
	return lists, nil
}
