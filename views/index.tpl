<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">

<html>
<head>
  {{ template "head.html" }}
</head>

<body>
  {{if .flash.error}}
  <blockquote>{{.flash.error}}</blockquote>
  {{end}}
  {{if .flash.notice}}
  <blockquote>{{.flash.notice}}</blockquote>
  {{end}}

  {{ template "header.html" . }}
  
  <div class="detail">
    {{ $isUserLogged := .IsUserLogged }}
    {{ $userRole := .UserRole }}
    {{ $userId := .UserId }}

    <div class="user-info">
    {{ if $isUserLogged }}
    Welcome back <b>{{ .UserName }}</b> | <a href="{{ urlfor "LoginController.Logout"}}" class="link-logout">Logout</a>
    {{ else }}
    <a href="{{ urlfor "LoginController.Get"}}" class="link-login">Login</a>
    {{ end }}
    </div>

    {{ if .Books }}
    <div class="books-container">
      <h2 class="books-title">Books Available</h2>
      <table class="books-table">
      <tr>
        <th>Title</th>
        <th>Isbn</th>
        <th>Author</th>
        <th>Available Copies</th>
        <th>Reviews</th>
        {{ if $isUserLogged }}
        <th></th>
        {{ end }}
      </tr>
      
      {{ range .Books }}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Isbn}}</td>
        <td>{{.Author}}</td>
        <td>{{.AvailableCopies}}</td>
        <td>
        {{ if eq (.GetReviews | len) 0 }}
          {{ if $isUserLogged }}
            <a href="{{ urlfor "UserBookReviewController.Get" "bookid" .Id}}">Add review</a>
          {{ else }}
            0
          {{ end }}
        {{ else }}
          <a href="{{ urlfor "UserBookReviewController.Reviews" "bookid" .Id}}">{{ .GetReviews | len }}</a>
        {{ end }}
        </td>
        {{ if $isUserLogged }}
        {{ if eq $userRole "adm" }}
        <td><a href="{{ urlfor "BookController.Get" "id" .Id }}">Edit</a> | <a href="{{ urlfor "BookController.Remove" "id" .Id }}">Delete</a></td>
        {{ else }}
        <td> 
          {{ if HasBook $userId .Id }}
          <a href="{{ urlfor "UserBookController.ReturnBook" "bookid" .Id}}">Return</a>
          {{ else }}
          <a href="{{ urlfor "UserBookController.LendBook" "bookid" .Id}}">Lend a copy</a>
          {{ end }}
        </td>
        {{ end }}
        {{ end }}
      </tr>
      {{ end }}
      </table>
    </div>
    {{ else }}
      <p class="no-books">No books available</p>
    {{ end }}

    {{ if $isUserLogged }}
    {{ if eq $userRole "adm" }}
    <div class="admin-actions">
      <a href="{{ urlfor "BookController.Get" }}" class="btn-add-book">Add new book</a>
    </div>
    {{ end }}
    {{ end }}
  </div>
  {{ template "footer.html" }}
</body>
</html>
