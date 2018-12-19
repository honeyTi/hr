{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="">我的博客</a>
        <ul class="nav navbar-nav">
            <li><a {{if .IsHome}}class="active"{{end}} href="/index">首页</a></li>
            <li><a {{if .Category}}class="active"{{end}} href="/category">分类</a></li>
            <li><a {{if .Topic}}class="active"{{end}} href="/topic">文章</a></li>
        </ul>
    </div>
</div>
<div class="pull-right">
    <ul class="nav navbar-nav">
    {{if .IsLogin}}
        <li><a href="/login?exit=true">退出</a></li>
    {{else}}
        <li><a href="/login">管理员登录</a></li>
    {{end}}
    </ul>
</div>
{{end}}