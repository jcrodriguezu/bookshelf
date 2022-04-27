# bookshelf
This is a demo project for Beego. 
It simulates a bookshelf where the user can lend, return, and review a book; the administrator role can create and delete books.

check the beego documentation in: https://beego.vip/docs/intro/

## Features:
* Controllers: Sessions, Flash messages and redirects, external api call to extract the book name and the author by the isbn
* Filters: Auth validation filter and redirect
* Forms: Generic form validation (using form tags and interfaces), form to model conversion
* Models: Orm tags, orm crud operations using sqlite3
* Routers: use default methods and custom methods, separation of routes by namespaces
* Tests: Index and login tests (with redirect validation) using convey
* Views: Use different template ext, render form, custom template functions

## Execution
First install and build:

```> go get```

```> go build```

Before executing the created binary, if is the first time executing this example, create an `.env` file and  then add the following variable to that file: 
`DB_INITIAL_FORCE=true`

execute the app

```> ./bookshelf```

After the initial execution you can set to false the variable `DB_INITIAL_FORCE=false` to avoid the initial data creation.

## Tests execution
Go to the tests folder

```> cd tests```

And execute the tests

```> go test```


## Licence
[GNU General Public License](https://www.gnu.org/licenses/gpl-3.0.en.html)
