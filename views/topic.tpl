<!DOCTYPE html>

<html>

{{template "header"}}
<title>文章-我的博客</title>
</head>
<body>
{{template "navbar" .}}
<div class="container" style="padding-top: 40px;">
    <h1>分类列表</h1>
    <a href="/topic/add" class="btn btn-default">添加文章</a>
    <table class="table table-bordered">
        <thead>
        <tr>
            <th>#</th>
            <th>文章名称</th>
            <th>最后更新</th>
            <th>浏览</th>
            <th>回复数</th>
            <th>操作</th>
        </tr>
        <tbody>
        {{range .Topics}}
        <tr>
            <td>{{.Id}}</td>
            <td>{{.Title}}</td>
            <td>{{.Updated}}</td>
            <td>{{.Views}}</td>
            <td>{{.ReplyCount}}</td>
            <td><a href="/topic/delete/{{.Id}}">删除</a></td>
        </tr>
        {{end}}
        </tbody>
        </thead>
    </table>
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>

</body>
</html>
