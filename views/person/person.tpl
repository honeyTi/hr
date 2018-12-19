<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <div>{{.userName}}</div>
    <h1>{{.Teb}}</h1>

    <h3>
        {{if .trueCode}} 内容为真 {{end}}
    </h3>
    <h3>
        {{if .falseCode}}
        {{else}}
        内容为假的
        {{end}}
    </h3>
    {{template "test"}}
    <div>
        {{with .user}}
                {{.Name}}
                {{.Age}}
         {{end}}
    </div>
    <div>
        {{range .nums}}
          {{.}}
        {{end}}
    </div>
    {{str2html .html}}

    <div>
        {{.Pipe | htmlquote}}
    </div>
</body>
</html>
{{define "test"}}
<div>
    this is template
</div>
{{end}}