{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="">我的博客</a>
        <ul class="nav navbar-nav">
            <li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
            <li  {{if .Category}}class="active"{{end}}><a href="/category">分类</a></li>
            <li {{if .Topic}}class="active"{{end}} ><a href="/topic">文章</a></li>
        </ul>
        <ul class="nav navbar-nav pull-right">
        {{if .IsLogin}}
            <li style="float: right;"><a href="/login?exit=true">退出</a></li>
        {{else}}
            <li style="float: right;"><a href="/login">管理员登录</a></li>
        {{end}}
        </ul>
    </div>
</div>

{{end}}