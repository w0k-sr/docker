package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Level struct {
	Id        int   `orm:"column(id);auto"`
	Level     int   `orm:"column(level);"`
	ZeroNum   int   `orm:"column(zero_num);"`
	DeleteFlg bool  `orm:"column(delete_flg);"`
}

func (t *Level) TableName() string {
	return "level"
}

func init() {
	orm.RegisterModel(new(Level))
}

func GetLevelById(id int) (v *Level, err error) {
	o := orm.NewOrm()
	v = &Level{Level: id}
	if err = o.QueryTable(new(Level)).Filter("Level", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err

}