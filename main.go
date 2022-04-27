package main

import (
	"bookshelf/filters"
	"bookshelf/models"
	_ "bookshelf/routers"
	"bookshelf/utils"
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

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
		utils.InitialData()
	}

	return err
}

func main() {
	DB_INITIAL_FORCE := os.Getenv("DB_INITIAL_FORCE") == "true"

	if err := startDb(DB_INITIAL_FORCE); err != nil {
		panic(err)
	}

	web.BeeApp.InsertFilter("/logout", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/book/*", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/review/get", web.BeforeRouter, filters.AuthFilter)
	web.BeeApp.InsertFilter("/review/new", web.BeforeRouter, filters.AuthFilter)

	web.AddFuncMap("HasBook", models.HasBook)

	web.Run()
}
