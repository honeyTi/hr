<!DOCTYPE html>

<html>

{{template "header"}}
<title>分类-我的博客</title>
</head>
<body>
{{template "navbar" .}}
<div class="container" style="padding-top: 40px;">
    <h1>分类列表</h1>
    <form method="GET" action="/category">
        <div class="form-group">
            <label>用户名</label>
            <input type="text" id="name" class="form-control" placeholder="请输入分类名称" name="name">
        </div>
        <input type="hidden" name="op" value="add">
        <button type="submit" class="btn btn-default" onclick="checkInput()">添加</button>
    </form>
    <table class="table table-bordered">
        <thead>
        <tr>
            <th>#</th>
            <th>名称</th>
            <th>文章数</th>
            <th>操作</th>
        </tr>
        <tbody>
        {{range .Categories}}
        <tr>
            <td>{{.Id}}</td>
            <td>{{.Title}}</td>
            <td>{{.TopicCount}}</td>
            <td><a href="/category?op=del&id={{.Id}}">删除</a></td>
        </tr>
        {{end}}
        </tbody>
        </thead>
    </table>
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
<script type="text/javascript">
    function checkInput() {
        var uname = $("#name").val()
        if (uname.length === 0 || uname === undefined) {
            alert("请输入账号");
            return false
        }
        return true
    }
</script>
</body>
</html>
