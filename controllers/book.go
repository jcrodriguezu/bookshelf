package controllers

import (
	"bookshelf/forms"
	"bookshelf/models"
	"encoding/json"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

// BookController ...
type BookController struct {
	web.Controller
}

// Get ...
func (c *BookController) Get() {
	fd := web.ReadFromRequest(&c.Controller)
	c.Data["flash"] = fd.Data

	bookform := &forms.BookForm{}
	c.Data["Action"] = "BookController.New"

	id, err := c.GetInt("id")
	if err == nil {
		book := &models.Book{Id: id}
		if err := book.Read(); err != nil {
			flash := web.NewFlash()
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			c.Data["Action"] = "BookController.Edit"
			bookform = &forms.BookForm{
				Id:     book.Id,
				Isbn:   book.Isbn,
				Title:  book.Title,
				Author: book.Author,
				Copies: book.Copies,
			}
		}
	}

	c.Data["Form"] = bookform
	c.TplName = "form.tpl"
}

// New book
func (c *BookController) New() {
	flash := web.NewFlash()

	bookForm := &forms.BookForm{}
	if err := c.ParseForm(bookForm); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 303)
	}

	book, err := forms.ToModel(bookForm)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := book.(*models.Book).Insert(); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Book successful created")
			flash.Store(&c.Controller)
		}
	}

	c.Redirect(c.URLFor("BookController.Get"), 303)
}

// Edit book
func (c *BookController) Edit() {
	user := c.GetSession("user")
	if user == nil {
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	flash := web.NewFlash()

	bookForm := &forms.BookForm{}
	if err := c.ParseForm(bookForm); err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor("MainController.Get"), 303)
	}

	book, err := forms.ToModel(bookForm)
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if err := book.(*models.Book).Update(); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		} else {
			flash.Notice("Book successful updated")
			flash.Store(&c.Controller)
		}
	}
	c.Redirect(c.URLFor("MainController.Get"), 303)
}

// Remove book
func (c *BookController) Remove() {
	user := c.GetSession("user")
	if user != nil {
		c.Redirect(c.URLFor("MainController.Get"), 307)
	}

	flash := web.NewFlash()

	id, err := c.GetInt("id")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		book := &models.Book{Id: id}
		if err := book.Delete(); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
		}
	}
	c.Redirect(c.URLFor("MainController.Get"), 303)
}

// Search the isbn number and fetch the book information
func (c *BookController) SearchIsbn() {
	if c.GetSession("user") == nil {
		return
	}

	isbn := c.Ctx.Input.Param(":isbn")

	// Clean ISBN: remove hyphens and spaces
	cleanISBN := strings.ReplaceAll(isbn, "-", "")
	cleanISBN = strings.ReplaceAll(cleanISBN, " ", "")

	c.Ctx.Output.Header("Content-Type", "application/json")

	// Use Open Library Books API
	apiURL := "https://openlibrary.org/api/books?bibkeys=ISBN:" + cleanISBN + "&format=json&jscmd=data"
	resp, err := http.Get(apiURL)
	if err != nil {
		c.Ctx.Output.Body([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.Output.Body([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		c.Ctx.Output.Body([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}

	// Check if we have any results (Open Library returns empty object if not found)
	if len(dat) == 0 {
		c.Ctx.Output.Body([]byte(`{"error": "No book found for this ISBN"}`))
		return
	}

	// Get the book data using the ISBN key
	bibKey := "ISBN:" + cleanISBN
	bookData, ok := dat[bibKey].(map[string]interface{})
	if !ok {
		c.Ctx.Output.Body([]byte(`{"error": "No book found for this ISBN"}`))
		return
	}

	// Extract title
	title := ""
	if titleVal, ok := bookData["title"].(string); ok {
		title = titleVal
	}

	// Extract author (Open Library returns array of author objects with "name" field)
	author := ""
	if authors, ok := bookData["authors"].([]interface{}); ok && len(authors) > 0 {
		if authorObj, ok := authors[0].(map[string]interface{}); ok {
			if authorName, ok := authorObj["name"].(string); ok {
				author = authorName
			}
		}
	}

	// Return the result
	d := map[string]string{"title": title, "author": author}
	enc_json, _ := json.Marshal(d)
	c.Ctx.Output.Body(enc_json)
}
