package controllers

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"math"
	"myapp/models"
	"myapp/common"
	beego "github.com/beego/beego/v2/server/web"
)

type SudokuController struct {
	beego.Controller
}

// 数独一覧表示
func (c *SudokuController) List() {
	data, err := models.GetSudokuGenerateWithDetails()
	message := c.GetSession("message")
	if err != nil {
		fmt.Println("Error:", err)
		c.Data["Message"] = "エラーが発生しました。管理者にお問い合わせください。"
	}
	if message != nil {
		c.Data["Message"] = message
		c.DelSession("message")
	}
	c.Data["SudokuGenerates"] = data
	c.TplName = "sudoku/list.tpl"
}

// 数独登録
func (c *SudokuController) Create() {
	var s models.SudokuGenerate
	var p models.Progress
	// 自動でレベルを生成
	rand.Seed(time.Now().UnixNano())
	s.Level = rand.Intn(3) + 1
	var cs common.Sudoku
	// 9×9の数字の配列生成
	cs= cs.AddSudoku()
	slice := cs.ToSlice()
	s.SudokuZeroAbsence, _  = common.ArrayToJson(slice)
	// 9×9の数字の配列からレベルごとの数だけ0へ返還
	data, _ := models.GetLevelById(s.Level)
	cs.RemoveDigits(data.ZeroNum)
	slice = cs.ToSlice()
	blanksudoku, _ := common.ArrayToJson(slice)
	s.SudokuZeroPresence = blanksudoku
	p.SudokuSolve = blanksudoku
	id, err := models.AddSudokuGenerate(&s)
	p.SudokuId = int(id)
	p.ProgressRate = 0
	_, err = models.AddProgress(&p)
	if err != nil {
		fmt.Println("Error:", err)
		c.SetSession("message", "エラーが発生しました。管理者にお問い合わせください。")
	}
	c.SetSession("message", "作成しました")
	c.Redirect("/sudoku/list", 302)
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

	preArray, _ := common.JsonToArray(data.SudokuZeroPresence)
	abArray, _ := common.JsonToArray(data.SudokuZeroAbsence)
	solveArray, _ := common.JsonToArray(data.SudokuSolve)
	// 解答チェック
	flagArray, _, comment := check(abArray, solveArray, data)
	c.Data["Sudoku"] = data
	c.Data["preArray"] = preArray
	c.Data["solveArray"] = solveArray
	c.Data["flagArray"] = flagArray
	c.Data["comment"] = comment
	c.TplName = "sudoku/detail.tpl"
}

// 数独の入力値を更新
func (c *SudokuController) Update() {
	// 画面から受け取ったjsonの値
	idNo, _ := c.GetInt("id")
	colNo, _ := c.GetInt("col")
	rowNo, _ := c.GetInt("row")
	numNo, _ := c.GetInt("num")
	data, _ := models.GetSudokuGenerateById(idNo)
	abArray, _ := common.JsonToArray(data.SudokuZeroAbsence)
	solveArray, _ := common.JsonToArray(data.SudokuSolve)
	flag := false
	if numNo != abArray[rowNo][colNo] && numNo != 0 {
		flag = true
	}
	solveArray[rowNo][colNo] = numNo
	zeroCount := 0
	for _, array := range solveArray {
		for _, num := range array {
			if num == 0 {
				zeroCount++
			}
		}
	}
	// 解答チェック
	_, count, comment := check(abArray, solveArray, data)
	result := zeroCount + count
	p, _ := models.GetProgressById(idNo)
	solveStr, _ := common.ArrayToJson(solveArray)
	p.SudokuSolve = solveStr
	p.ProgressRate = int(math.Ceil((float64(data.ZeroNum - result) / float64(data.ZeroNum)) * 100))
	models.UpdateProgressById(p)
	// JSONデータ
	response := map[string]interface{}{
		"status": "success",
		"message": "Update successful",
		"data": map[string]interface{}{
			"flag": flag,
			"comment": comment,
		},
	}

	// JSONレスポンスを返す
	c.Data["json"] = response
	c.ServeJSON()
	
	// 以降の処理を停止して画面遷移を防止
	c.StopRun()
}

// 解答チェック
func check(abArray [][]int, solveArray [][]int, v *models.SudokuGenerateWithDetails) ([][]bool, int, int) {
	var flagArray [][]bool
	var commentFlg bool
	count := 0
	// 正解 false 不正解 true
	for  j, array := range abArray {
		var f []bool
		for k, num := range array {
			flag := false
			if solveArray[j][k] != num && solveArray[j][k] != 0 {
				flag = true
				commentFlg = true
				// 不正解の数を数える
				count++
			}
			f = append(f, flag)
		}
		flagArray = append(flagArray, f)
	}
  // 詳細画面の画像とコメント判定
	comment := 1
	if commentFlg {
		comment = 2
	} else {
		if v.ProgressRate > 60 {
			comment = 3
		} else if v.ProgressRate == 100 {
			comment = 4
		}
	}
	return flagArray, count, comment
}