<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
    {{ template "head.html" }}
</head>
<body>
    {{ template "header.html" }}
    <div class="detail">
        <a href="index">Return</a><br><br>
        <form action="{{ urlfor "LoginController.Post" }}" method="post">
            username:
            <input type="text" name="username" value="" /><br />
            password:
            <input type="password" name="password" value="" /><br />

            <input type="submit" value="Log In" />
        </form>
    </div>
    {{ template "footer.html" }}
</body>
</html>