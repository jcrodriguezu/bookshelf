# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Bookshelf is a demo web application built with Beego v2.0.2 framework that simulates a library system where users can lend, return, and review books, while administrators can create and delete books.

## Build and Run Commands

**Initial setup:**
```bash
go get
go build
```

**First-time execution:**
Create a `.env` file with:
```
DB_INITIAL_FORCE=true
```

Then run:
```bash
./bookshelf
```

After the first run, change `.env` to `DB_INITIAL_FORCE=false` to prevent reinitializing the database.

**Running tests:**
```bash
cd tests
go test
```

**Development server:**
The application runs on `http://localhost:8080` by default (configured in `conf/app.conf`).

## Architecture

### MVC Pattern with Beego Framework

This application follows a strict MVC architecture with additional layers:

**Controllers** → **Forms** → **Models** → **Database**

### Key Architectural Patterns

#### 1. Form-to-Model Conversion Pattern
The codebase uses an interface-based pattern to convert and validate forms before persisting to models:

- `forms/iform.go` defines `IForm` interface with `ToModel()` method
- All form structs implement `IForm` and handle validation using `govalidator` tags
- `forms.ToModel()` validates the form struct, then calls the specific form's `ToModel()` implementation
- Controllers parse forms, validate via `forms.ToModel()`, then call model CRUD methods

**Example flow:**
```
BookController.New() → ParseForm(BookForm) → forms.ToModel(bookForm) → book.Insert()
```

#### 2. Session-Based Authentication
- Session management enabled globally in `conf/app.conf` (`SessionOn = true`)
- User object stored in session upon successful login
- `filters/auth.go` contains `AuthFilter` that redirects unauthenticated users
- Filters registered in `main.go` for protected routes (`/book/*`, `/review/*`, `/logout`)

#### 3. Namespace-Based Routing
Routes are organized using Beego namespaces in `routers/router.go`:
- `bookNs` namespace: `/book/*` routes
- `reviewNs` namespace: `/review/*` routes
- Direct routes for public pages (`/`, `/index`, `/login`)
- Special route for ISBN lookup: `/scrap/:isbn`

#### 4. ORM Relationship Management
Models use Beego ORM with explicit relationship loading:

**User ↔ Book (Many-to-Many):**
```go
User.BooksLent []*Book `orm:"rel(m2m)"`
```
Managed via `QueryM2M()` for lending/returning books.

**User → Review (One-to-Many):**
```go
User.Reviews []*Review `orm:"reverse(many)"`
Review.User *User `orm:"rel(fk)"`
```

**User → Role (Many-to-One):**
```go
User.Role *Role `orm:"null;rel(one);on_delete(set_null)"`
```

Relations must be explicitly loaded with `LoadRelated()` after reading models.

#### 5. Database Initialization Pattern
`main.go` uses environment-driven database initialization:
- `DB_INITIAL_FORCE=true`: Drops tables, recreates schema, seeds initial data via `utils.InitialData()`
- `DB_INITIAL_FORCE=false`: Connects to existing database without modifications
- Uses SQLite3 with file `bookshelf.db`

### External API Integration

**Open Library Books API:**
- Endpoint: `/scrap/:isbn` → `BookController.SearchIsbn()`
- API URL: `https://openlibrary.org/api/books?bibkeys=ISBN:{isbn}&format=json&jscmd=data`
- Cleans ISBN input (removes hyphens/spaces) before API call
- Returns JSON: `{"title": "...", "author": "..."}`
- Frontend JavaScript in `views/form.tpl` auto-fills title/author fields on ISBN blur event

**Response parsing:**
```go
// API returns: {"ISBN:9780385472579": {"title": "...", "authors": [{"name": "..."}]}}
bookData := dat["ISBN:" + cleanISBN]
title := bookData["title"].(string)
author := bookData["authors"][0]["name"].(string)
```

### Template System

Templates use Beego's rendering engine with custom functions:
- Template files: `views/*.tpl` (not `.html`)
- Custom function registered in `main.go`: `web.AddFuncMap("HasBook", models.HasBook)`
- Flash messages use Post-Redirect-Get pattern for user notifications
- Shared templates: `head.html`, `header.html`, `footer.html`

## Important Technical Constraints

### Security Warning
**Password storage is in plaintext** (see `models/user.go:49`). The code has a TODO noting this. When implementing authentication improvements, passwords should be hashed using bcrypt or similar.

### Form Validation
All forms use `govalidator` struct tags. The validation happens in `forms.ToModel()` before any database operations.

### ORM Notes
- Models must call `orm.RegisterModel()` in their `init()` function
- Always use `LoadRelated()` to fetch relationship data after `Read()`
- M2M relationships require `QueryM2M()` API for add/remove operations
- Delete operations on referenced models respect `on_delete` ORM tags

### Filter Execution Order
Filters in `main.go` execute at `BeforeRouter` stage, before route matching. Protected routes require active session or redirect to `/`.

### Testing
Tests use `smartystreets/goconvey` framework (see `tests/default_test.go` for examples of testing redirects and session behavior).
