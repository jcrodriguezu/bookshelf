<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
    {{ template "head.html" }}
</head>
<body>
    {{ template "header.html" }}
    <div class="detail">
        <a href="index">Return</a><br><br>
        <form action="{{ urlfor "BookController.Post" }}" method="post">
            Title:
            <input type="text" name="title" value="" /><br />
            Author:
            <input type="text" name="author" value="" /><br />
            # Copies:
            <input type="text" name="copies" value="" /><br />

            <input type="submit" value="Create" />
        </form>
    </div>
    {{ template "footer.html" }}
</body>
</html>