package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type Progress struct {
	ProgressId   int    `orm:"column(progress_id);auto"`
	SudokuId     int    `orm:"column(sudoku_id);"`
	ProgressRate int    `orm:"column(progress_rate);"`
	SudokuSolve  string `orm:"column(sudoku_solve);"`
	DeleteFlg    bool   `orm:"column(delete_flg);"`
}

func init() {
	orm.RegisterModel(new(Progress))
}

func AddProgress(m *Progress) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err != nil {
		fmt.Println("Progress:", err)
	}
	return
}

func GetProgressById(id int) (v *Progress, err error) {
	o := orm.NewOrm()
	v = &Progress{SudokuId: id}
	if err = o.QueryTable(new(Progress)).Filter("SudokuId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	fmt.Println("Progress:", err)
	return nil, err
}

func UpdateProgressById(m *Progress) (err error) {
	o := orm.NewOrm()
	v := Progress{ProgressId: m.ProgressId}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}