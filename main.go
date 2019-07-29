package main

import (
	"fmt"
	_ "studentgo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.BConfig.CopyRequestBody = true

	//Get database
	dbUser := beego.AppConfig.String("mysqluser")
	dbPwd := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	dbString := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName)

	//Register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	//Register default database
	orm.RegisterDataBase("default", "mysql", dbString)

	//Autosync
	//db alias
	name := "default"

	//drop table and re-create
	force := false

	//print log
	verbose := true

	//error
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	//Run beego
	beego.Run()
}
