package main

import (
	"bookshelf/filters"
	"bookshelf/models"
	_ "bookshelf/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase("default", "sqlite3", "file:bookshelf.db")
}

func startDb(force bool) error {
	// Database alias.
	name := "default"
	// Print log.
	verbose := true
	// Debug
	orm.Debug = false
	// Error.
	err := orm.RunSyncdb(name, force, verbose)

	if force {
		initialData()
	}

	return err
}

func initialData() {
	o := orm.NewOrm()

	ra := new(models.Role)
	ra.Name = "adm"
	fmt.Println(o.Insert(ra))

	u1 := new(models.User)
	u1.Name = "Admin"
	u1.Username = "admin"
	u1.Password = "admin"
	u1.Role = ra
	fmt.Println(o.Insert(u1))

	ru := new(models.Role)
	ru.Name = "usr"
	fmt.Println(o.Insert(ru))

	u2 := new(models.User)
	u2.Name = "test"
	u2.Username = "test"
	u2.Password = "test"
	u2.Role = ru
	fmt.Println(o.Insert(u2))
}

func main() {
	if err := startDb(false); err != nil {
		panic(err)
	}

	web.BeeApp.InsertFilter("/logout", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/book/*", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/review/get", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/review/new", web.BeforeRouter, filters.AuthFilter)

	web.AddFuncMap("HasBook", models.HasBook)

	web.Run()
}
