/**
  * Created by: Jianyi
  * Date: 2019/1/4
  * Time: 13:37
  * Description:
 **/
package models

import (
	"github.com/astaxie/beego/orm"
)

type Leader struct{
	Id         int       `pk:"auto"`
	Name      string    `orm:"size(40)"`
	ProField string `orm:"size(200)"`
	Wechat string `orm:"size(100)"`
	Status     int `orm:"size(11)"`
	Avatar	string `orm:"size(255)"`
	WechatImg	string `orm:"size(255)"`
}

func (this *Leader) TableName() string {
	return TableName("leader")
}

func LeaderAdd(record *Leader) (int64, error) {
	return orm.NewOrm().Insert(record)
}

func (this *Leader) LeaderUpdate(fields ... string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func LeaderGetList(page, pageSize int, filters ...interface{}) ([]*Leader, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Leader, 0)
	query := orm.NewOrm().QueryTable(TableName("leader"))
	if len(filters) > 0 && filters != nil {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

//只显示正常状态
func LeaderAll() []*Leader {
	list := make([]*Leader, 0)
	orm.NewOrm().QueryTable(TableName("leader")).Filter("status",1).All(&list)
	return list
}

func LeaderGetById(id int) (*Leader, error) {
	leader := new(Leader)
	err := orm.NewOrm().QueryTable(TableName("leader")).Filter("id", id).One(leader)
	if err != nil {
		return nil, err
	}
	return leader, nil
}
