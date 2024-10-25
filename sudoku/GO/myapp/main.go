package main

import (
	// "strconv"
	"fmt"
	"os"
	_ "myapp/routers"
	// "github.com/astaxie/beego"
	beego "github.com/beego/beego/v2/server/web"
	// beeLogger "github.com/beego/bee/v2/logger"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	err := godotenv.Load(".env")
	
	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		// beeLogger.Log.Fatal("読み込み出来ませんでした: %s" + err.Error())
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	sqlConn := os.Getenv("SQL_CON")
	fmt.Println(sqlConn)
	fmt.Println("%s",sqlConn)
	orm.RegisterDataBase("default", "postgres", sqlConn)
	beego.Run()
}

