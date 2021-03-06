package models

import "github.com/astaxie/beego/orm"

type Contract struct {
	Id     int    `pk:"auto"`
	Content   string `orm:"size(255)"`
	Status int8   `orm:"size(1)"`
}

func (this *Contract) TableName() string {
	return TableName("contract")
}

func ContractAdd(record *Contract) (int64, error) {
	return orm.NewOrm().Insert(record)
}

func (this *Contract) ContractUpdate(fields ... string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func ContractGetList(page, pageSize int, filters ...interface{}) ([]*Contract, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Contract, 0)
	query := orm.NewOrm().QueryTable(TableName("contract"))
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

func ContractGetListApi(page, pageSize int, filters ...interface{}) ([]*Contract, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Contract, 0)
	query := orm.NewOrm().QueryTable(TableName("contract"))
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

func ContractGetById(id int) (*Contract, error) {
	Contract := new(Contract)
	err := orm.NewOrm().QueryTable(Contract.TableName()).Filter("id", id).One(Contract)
	if err != nil {
		return nil, err
	}
	return Contract, nil
}
