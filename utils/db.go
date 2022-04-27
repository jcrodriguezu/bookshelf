package utils

import (
	"bookshelf/models"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

func InitialData() {
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
