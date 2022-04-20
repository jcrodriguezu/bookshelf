package tests

import (
	"bookshelf/models"
	_ "bookshelf/routers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/smartystreets/goconvey/convey"
)

func HasBook() bool {
	return false
}

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	os.Setenv("BEEGO_RUNMODE", "test")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:test.db")

	web.AddFuncMap("HasBook", HasBook)
	web.TestBeegoInit(apppath)
}

// TestDefault is a sample to run an endpoint test
func TestDefault(t *testing.T) {
	Convey("Subject: Test Index\n", t, func() {
		Convey("No Books\n", func() {
			orm.RunSyncdb("default", true, false)
			r, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("The login button Should Exist", func() {
				So(w.Body.String(), ShouldContainSubstring, "<a href=\"/login\">Login</a>")
			})
			Convey("No Books Should Be Available", func() {
				So(w.Body.String(), ShouldContainSubstring, "No books available")
			})
		})
		Convey("One Book\n", func() {
			orm.RunSyncdb("default", true, false)
			o := orm.NewOrm()
			book := &models.Book{Title: "test", Author: "test", Copies: 1}
			o.Insert(book)

			r, _ := http.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
			Convey("The login button Should Exist", func() {
				So(w.Body.String(), ShouldContainSubstring, "<a href=\"/login\">Login</a>")
			})
			Convey("Should Have Books Available", func() {
				So(w.Body.String(), ShouldContainSubstring, "Books available")
			})
		})
	})
}

func TestLogin(t *testing.T) {
	Convey("Subject: Test Login\n", t, func() {
		orm.RunSyncdb("default", true, false)
		o := orm.NewOrm()
		role := &models.Role{Name: "usr"}
		o.Insert(role)
		u := &models.User{Name: "Test", Username: "test", Password: "test", Role: role}
		o.Insert(u)

		Convey("Wrong\n", func() {
			r, _ := http.NewRequest("POST", "/login", nil)
			r.Form = url.Values{
				"username": {"wrong"},
				"password": {"test"},
			}
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status Code Should Be 303\t", func() {
				So(w.Code, ShouldEqual, 303)
			})
			Convey("Should redirect to /login\t", func() {
				// NOTE not sure if there is other way to test the redirect
				So(w.HeaderMap["Location"], ShouldContain, "/login")
			})
			Convey("Should Have Error\t", func() {
				So(len(w.HeaderMap["Set-Cookie"]), ShouldEqual, 2)
				So(w.HeaderMap["Set-Cookie"][1], ShouldContainSubstring, "Wrong+Username+or+Password")
			})
		})
		Convey("Correct\n", func() {
			r, _ := http.NewRequest("POST", "/login", nil)
			r.Form = url.Values{
				"username": {"test"},
				"password": {"test"},
			}
			w := httptest.NewRecorder()
			web.BeeApp.Handlers.ServeHTTP(w, r)

			Convey("Status Code Should Be 303\t", func() {
				So(w.Code, ShouldEqual, 303)
			})
			Convey("Should redirect to /index\t", func() {
				// NOTE not sure if there is other way to test the redirect
				So(w.HeaderMap["Location"], ShouldContain, "/index")
			})
			Convey("Should Not Have Error\t", func() {
				So(len(w.HeaderMap["Set-Cookie"]), ShouldEqual, 1)
			})
		})
	})

}
