<!DOCTYPE html>

<html>

{{template "header"}}
<title>修改文章-我的博客</title>
</head>
<body>
{{template "navbar" .}}
<div class="container" style="padding-top: 40px;">
    <h1>修改文章</h1>
    <form action="/topic" method="post">
        <input type="hidden" name="tid" value="{{.Tid}}">
        <div class="form-group">
            <label for="">文章标题：</label>
            <input type="text" name="title" class="form-control" value="{{.Topic.Title}}">
        </div>
        <div class="form-group">
            <label for="">文章内容：</label>
            <textarea name="content" class="form-control" cols="30" rows="10">
                {{.Topic.Content}}
            </textarea>
        </div>
        <button type="submit" class="btn btn-default">修改文章</button>
    </form>
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
<script>

</script>
</body>
</html>
