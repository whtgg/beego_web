package models

import (
	"github.com/astaxie/beego/orm"
)

type Banner struct {
	Id     int    `pk:"auto"`
	Path   string `orm:"size(255)"`
	Status int8   `orm:"size(1)"`
	Link   string `orm:"size(255)"`
}

func (this *Banner) TableName() string {
	return TableName("banner")
}

func BannerAdd(record *Banner) (int64, error) {
	return orm.NewOrm().Insert(record)
}

func (this *Banner) BannerUpdate(fields ... string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func BannerGetList(page, pageSize int, filters ...interface{}) ([]*Banner, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Banner, 0)
	query := orm.NewOrm().QueryTable(TableName("banner"))
	query = query.Filter("status__in", 1, 5)
	if len(filters) > 0 && filters != nil {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

func BannerGetById(id int) (*Banner, error) {
	banner := new(Banner)
	err := orm.NewOrm().QueryTable(TableName("banner")).Filter("id", id).One(banner)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

func BannerGetListApi(page, pageSize int, filters ...interface{}) ([]*Banner, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Banner, 0)
	query := orm.NewOrm().QueryTable(TableName("banner"))
	query = query.Filter("status__in", 1)
	if len(filters) > 0 && filters != nil {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}
