{{template "header"}}
<title>我的博客-登录</title>
</head>
<html>
<body>
<div class="container" style="width:500px;margin-top: 100px">
    <form method="post" action="/login">
        <div class="form-group">
            <label>用户名</label>
            <input type="text" id="uname" class="form-control" placeholder="请输入用户名" name="uname">
        </div>
        <div class="form-group">
            <label>密码</label>
            <input type="password" id="pwd" class="form-control"placeholder="请输入密码" name="pwd">
        </div>
        <div class="checkbox">
            <label>
                <input type="checkbox" name="autoLogin"> 自动登录
            </label>
        </div>
        <button type="submit" class="btn btn-default" onclick="checkInput()">登录</button>
    </form>
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
<script type="text/javascript">
    function checkInput() {
        var uname = $("#uname").val()
        if (uname.length === 0 || uname === undefined) {
            alert("请输入账号");
            return false
        }
        var uname = $("#pwd").val()
        if (uname.length === 0 || uname === undefined) {
            alert("请输入密码");
            return false
        }
        return true
    }
</script>
</body>
</html>