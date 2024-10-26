package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type SudokuGenerate struct {
	Id                 int     `orm:"column(id);auto"`
	Level              int     `orm:"column(sudoku_level);"`
	SudokuZeroPresence string  `orm:"column(sudoku_zero_presence);null;`
	SudokuZeroAbsence  string  `orm:"column(sudoku_zero_absence);null;`
	DeleteFlg          bool    `orm:"column(delete_flg);"`
}

type SudokuGenerateWithDetails struct {
	SudokuGenerate
	Progress
	Level	
}

func (t *SudokuGenerate) TableName() string {
	return "sudoku_generate"
}

func init() {
	orm.RegisterModel(new(SudokuGenerate))
}

func AddSudokuGenerate(m *SudokuGenerate) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err != nil {
		fmt.Println("SudokuGenerate:", err)
	}
	return
}

func GetSudokuGenerateWithDetails()([]SudokuGenerateWithDetails, error) {
		o := orm.NewOrm()
		var results []SudokuGenerateWithDetails
		sql := `SELECT 
		s.id,
    s.sudoku_level,
    s.sudoku_zero_presence,
    s.sudoku_zero_absence,
		l.zero_num,
		p.progress_rate,
		p.sudoku_solve
		FROM sudoku_generate s 
		LEFT JOIN progress p ON s.id = p.sudoku_id
		INNER JOIN level l ON s.sudoku_level = l.level
		ORDER BY s.id`
    _, err := o.Raw(sql).QueryRows(&results)
    if err != nil {
        return nil, err
				fmt.Println("GetSudokuGenerateWithDetails:", err)
    }

    return results, nil
}

func GetSudokuGenerateById(id int) (v *SudokuGenerateWithDetails, err error) {
	o := orm.NewOrm()
	result := &SudokuGenerateWithDetails{}
	sql := `SELECT 
	s.id,
	s.sudoku_level,
	s.sudoku_zero_presence,
	s.sudoku_zero_absence,
	l.zero_num,
	p.progress_id,
	p.progress_rate,
	p.sudoku_solve
	FROM sudoku_generate s 
	INNER JOIN progress p ON s.id = p.sudoku_id
	INNER JOIN level l ON s.sudoku_level = l.level
	WHERE s.id = ?
	ORDER BY s.id
	`
	err = o.Raw(sql, id).QueryRow(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}