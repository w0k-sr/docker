package routers

import (
	"myapp/controllers"
	// "github.com/astaxie/beego"
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
		)
	beego.AddNamespace(ns)
}
