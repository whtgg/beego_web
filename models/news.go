package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type News struct {
	Id         int       `pk:"auto"`
	Title      string    `orm:"size(100)"`
	SubTitle   string    `orm:"size(100)"`
	Content    string
	CreateTime time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdateTime time.Time `orm:"auto_now;type(timestamp)"`
	Publish    string
	Status     int8
	Weight     int
	Cover      string
	Genre      int8
	Tags	   int
}

func (this *News) TableName() (names string) {
	names = TableName("adm_news")
	return
}

func NewsAdd(this *News) (int64, error) {
	return orm.NewOrm().Insert(this)
}

func (self *News) NewsUpdate(fields ... string) error {
	if _, err := orm.NewOrm().Update(self, fields...); err != nil {
		return err
	}
	return nil
}

func NewsGetList(page, pageSize int, filters ...interface{}) ([]*News, int64) {
	offset := (page - 1) * pageSize
	list := make([]*News, 0)
	query := orm.NewOrm().QueryTable(TableName("adm_news"))
	if len(filters) > 0 && filters != nil {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-create_time").Limit(pageSize,offset).All(&list)
	return list, total
}

func NewsGetTrend(page, pageSize int) ([]*News, int64) {
	o := orm.NewOrm()
	offset := (page - 1) * pageSize
	list := make([]*News, 0)
	_, err := o.Raw("SELECT * FROM pp_adm_news WHERE status=1 order by create_time desc limit ? offset ? ", pageSize, offset).QueryRows(&list)
	count := make([]*News, 0)
	if err == nil {
		total, _ := o.Raw("SELECT id FROM pp_adm_news WHERE status=1").QueryRows(&count)
		return list, total
	}
	return list, 0
}

func NewsGetById(id int) (*News, error) {
	news := new(News)
	err := orm.NewOrm().QueryTable(TableName("adm_news")).Filter("id", id).One(news)
	if err != nil {
		return nil, err
	}
	return news, nil
}

func NewsGetByIdRange(id int)(*News,*News,error){
	prev := new(News)
	next := new(News)
	err := orm.NewOrm().QueryTable(TableName("adm_news")).Filter("id__lt", id).OrderBy("-id").One(prev,"Id","Title")
	err = orm.NewOrm().QueryTable(TableName("adm_news")).Filter("id__gt", id).One(next,"Id","Title")
	return prev,next,err
}
