<!DOCTYPE html>

<html>

{{template "header"}}
<title>首页-我的博客</title>
</head>
<body>
{{template "navbar" .}}
<div class="container" style="padding-top: 40px;">
{{range .Topics}}
    <div class="page-header">
        <h1>{{.Title}}</h1>
        <h6 class="text-muted">
            发表于{{.Created}} 浏览量{{.Views}} 评论{{.ReplyCount}}
        </h6>
        <p>
            {{.Content}}
        </p>
    </div>
{{end}}
</div>
<script type="text/javascript" src="https://cdn.staticfile.org/jquery/3.3.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/bootstrap-3.3.7-dist/js/bootstrap.min.js"></script>
</body>
</html>
