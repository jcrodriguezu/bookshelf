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
    <div>
    {{ if $isUserLogged }}
    Welcome back <b>{{ .UserName }}</b> | <a href="{{ urlfor "LoginController.Logout"}}">Logout</a>
    {{ else }}
    <a href="{{ urlfor "LoginController.Get"}}">Login</a>
    {{ end }}
    </div>
    {{ if .Books }}
    <table border="1" style="width:100%">
      <caption>Books available</caption>
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
    {{ else }}
      No books available
    {{ end }}
  </div>
  <div>
  {{ if $isUserLogged }}
  {{ if eq $userRole "adm" }}
  <a href="{{ urlfor "BookController.Get" }}">Add new book</a>
  {{ end }}
  {{ end }}
  </div>
  {{ template "footer.html" }}
</body>
</html>
