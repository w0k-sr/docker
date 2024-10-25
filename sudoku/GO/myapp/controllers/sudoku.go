package controllers

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"myapp/models"
	"myapp/common"
	// "github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
)

type SudokuController struct {
	beego.Controller
}

// 数独一覧表示
func (c *SudokuController) List() {
	message := c.GetSession("message")
	if message != nil {
		c.Data["Message"] = message
		c.DelSession("message")
	}
	data, _ := models.GetSudokuGenerateWithDetails()
	c.Data["SudokuGenerates"] = data
	// fmt.Println(data)
	c.TplName = "sudoku/list.tpl"
}

// 数独登録
func (c *SudokuController) Create() {
	var s models.SudokuGenerate
	var p models.Progress
	rand.Seed(time.Now().UnixNano())
	s.Level = rand.Intn(3) + 1
	data, _ := models.GetLevelById(s.Level)
	var cs common.Sudoku
	cs= cs.AddSudoku()
	fmt.Println(data)
	// 配列からスライスに変換
	slice := cs.ToSlice()
	s.SudokuZeroAbsence, _  = common.ArrayToJson(slice)
	fmt.Println(s.SudokuZeroPresence)
	cs.RemoveDigits(data.ZeroNum)
	slice = cs.ToSlice()
	// var blanksudoku string
	blanksudoku, _ := common.ArrayToJson(slice)
	s.SudokuZeroPresence = blanksudoku
	p.SudokuSolve = blanksudoku
	id, _ := models.AddSudokuGenerate(&s)
	p.SudokuId = int(id)
	p.ProgressRate = 0
	models.AddProgress(&p)
	// fmt.Println("挿入されたID:", s.ID)
	c.SetSession("message", "作成しました")
	c.Redirect("/sudoku/list", 302)
}

// @router / [post]
func (c *SudokuController) Post() {

}

// 数独詳細
func (c *SudokuController) Get() {
	var i int
	idStr := c.Ctx.Input.Param(":id")
	i, _ = strconv.Atoi(idStr)
	data, err := models.GetSudokuGenerateById(i)
	if err != nil {
    fmt.Println("Error:", err)
		c.SetSession("message", "エラーが発生しました。管理者にお問い合わせください。")
		c.Redirect("/sudoku/list", 302)
	}

	if data == nil {
			fmt.Println("No data found for id:", i)
			c.SetSession("message", "数独が存在しません。")
			c.Redirect("/sudoku/list", 302)
	}

	// データが存在する場合の処理
	preArray, _ := common.JsonToArray(data.SudokuZeroPresence)
	// abArray, _ := common.ParsePostgresArray(data.SudokuZeroAbsence)
	solveArray, _ := common.JsonToArray(data.SudokuSolve)
	c.Data["Sudoku"] = data
	c.Data["preArray"] = preArray
	// c.Data["abArray"] = abArray
	c.Data["solveArray"] = solveArray
	// fmt.Println(solveArray)
	c.TplName = "sudoku/detail.tpl"
}

