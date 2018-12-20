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
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
</body>
</html>
