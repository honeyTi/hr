<!DOCTYPE html>

<html>

{{template "header"}}
<title>{{.Topic.Title}}-我的博客</title>
</head>
<body>
{{template "navbar" .}}
<div class="container" style="padding-top: 40px;">
    <div class="page-header">
        <h1>
        {{.Topic.Title}}
            <a href="/topic/modify?tid={{.Tid}}" class="pull-right btn btn-default">修改文章</a>
        </h1>
        <p>
        {{.Topic.Content}}
        </p>
    </div>
</div>
<div class="container">
    <h3>本文回复</h3>
    <form action="/replay" method="add">
        <input type="hidden" id="tid" value="{{.Topic.Id}}">
        <div class="form-group">
            <label>显示昵称：</label>
            <input type="text" class="form-control" name="nickname">
        </div>
        <div class="form-group">
            <label>内容：</label>
            <textarea name="content" id="" cols="30" rows="10" class="form-control"></textarea>
        </div>
        <button type="submit" class="btn btn-default">提交回复</button>
    </form>
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
</body>
</html>
