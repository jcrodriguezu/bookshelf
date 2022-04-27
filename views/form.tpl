<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
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

    {{ template "header.html" }}

    <div class="detail">
        <a href="/index">Return</a><br><br>
        <form action="{{ urlfor .Action }}" method="post">
            {{.Form | renderform}}
            <br />
            <input type="submit" value="Save" />
        </form>
    </div>
    {{ template "footer.html" }}

<script>
document.body.onload = function(){
    document.getElementById("isbn").onblur = function() {
        console.log("Lost focus");
        const isbn = document.getElementById("isbn").value;
        if (isbn === "") {
            console.log("Isbn is empty");
        } else {
            searchBook(isbn)
        }
    }
}


function searchBook(book_isbn) {
    const xhttp = new XMLHttpRequest();
    xhttp.onload = function() {
        const resp = JSON.parse(this.responseText)
        if (resp) {
            document.getElementById("title").value = resp.title;
            document.getElementById("author").value = resp.author;
        }
    }

    xhttp.open("GET", "/scrap/" + book_isbn, true);
    xhttp.send();
}
</script>

</body>
</html>