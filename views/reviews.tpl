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
  
  <a href="{{ urlfor "UserBookReviewController.Get" "bookid" .Id }}">New review</a> | <a href="/index">Return</a><br><br>
  <div class="review">
  {{ range .Reviews }}
  <div class="review-detail">
    <div class="review-title">{{ .Title }}</div>
    <div class="review-body">{{ .Body }}</div>
  </div>
  {{ end }}
  </div>

  {{ template "footer.html" }}
</body>
</html>
