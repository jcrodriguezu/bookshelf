package main

import (
	"bookshelf/models"
	_ "bookshelf/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)

	orm.RegisterDataBase("default", "sqlite3", "file:bookshelf.db")
}

func startDb() error {
	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true
	// Debug
	orm.Debug = true
	// Error.
	err := orm.RunSyncdb(name, force, verbose)

	if force {
		initialData()
	}

	return err
}

func initialData() {
	o := orm.NewOrm()
	o.Using("default")

	r := new(models.Role)
	r.Name = "usr"
	fmt.Println(o.Insert(r))

	u := new(models.User)
	u.Name = "Juanca"
	u.Username = "juanca"
	u.Password = "juanca"
	u.Role = r
	fmt.Println(o.Insert(u))

	b := new(models.Book)
	b.Author = "Popeye"
	b.Copies = 3
	b.Title = "The new adventures of Popeye"
	fmt.Println(o.Insert(b))
}

func main() {
	if err := startDb(); err != nil {
		panic(err)
	}

	beego.Run()
}
