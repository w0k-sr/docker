package routers

import (
	"myapp/controllers"
	beego "github.com/beego/beego/v2/server/web"

)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/sudoku",
		beego.NSRouter("/list", 
			&controllers.SudokuController{}, "get:List"),
		beego.NSRouter("/:id", 
			&controllers.SudokuController{}, "get:Get"),
		beego.NSRouter("/create", 
			&controllers.SudokuController{}, "get:Create"),
		beego.NSRouter("/update", 
			&controllers.SudokuController{}, "post:Update"),
		)
	beego.AddNamespace(ns)
}
