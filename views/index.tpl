<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">

<html>
<head>
  {{ template "head.html" }}
</head>

<body>
  {{ template "header.html" . }}
  <div class="detail">
    {{ $isUserLogged := .IsUserLogged }}
    {{ $userRole := .UserRole }}
    <div>
    {{ if $isUserLogged }}
    Welcome back {{ .UserName }} | <a href="{{ urlfor "LoginController.Logout"}}">Logout</a>
    {{ else }}
    <a href="{{ urlfor "LoginController.Get"}}">Login</a>
    {{ end }}
    </div>
    {{ if .Books }}
    <table border="1" style="width:100%">
      <caption>Books available</caption>
      <tr>
        <th>Title</th>
        <th>Author</th>
        <th>Available Copies</th>
        {{ if $isUserLogged }}
        <th></th>
        {{ end }}
      </tr>
      
      {{ range .Books }}
      <tr>
        <td>{{.Title}}</td>
        <td>{{.Author}}</td>
        <td>{{.AvailableCopies}}</td>
        {{ if $isUserLogged }}
        {{ if eq $userRole "adm" }}
        <td>Edit | Delete</td>
        {{ else if eq $userRole "usr" }}
        <td>Review</td>
        {{ end }}
        {{ end }}
      </tr>
      {{ end }}
    </table>
    {{ else }}
      No books available
    {{ end }}
  </div>
  {{ template "footer.html" }}
</body>
</html>
