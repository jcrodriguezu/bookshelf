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
	o.Using("default")

	ra := new(models.Role)
	ra.Name = "adm"
	fmt.Println(o.Insert(ra))

	u1 := new(models.User)
	u1.Name = "Juanca"
	u1.Username = "juanca"
	u1.Password = "juanca"
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

	b := new(models.Book)
	b.Author = "Popeye"
	b.Copies = 3
	b.Title = "The new adventures of Popeye"
	fmt.Println(o.Insert(b))
}

func main() {
	if err := startDb(false); err != nil {
		panic(err)
	}

	beego.AddFuncMap("HasBook", models.HasBook)

	beego.Run()
}
