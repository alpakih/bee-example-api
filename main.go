package main

import (
	_ "bee-example-api/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@localhost:5433/beego_api?sslmode=disable")
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//// Database alias.
	//name := "default"
	//
	//// Error.
	//err := orm.RunSyncdb(name, true, true)
	//if err != nil {
	//	fmt.Println(err)
	//}

	beego.Run()
}
