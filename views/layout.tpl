<!DOCTYPE html>
<html data-theme="light">
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="{{.BaseUrl}}/static/css/output.css">
    <link rel="icon" type="image/x-icon" href="/favicon.ico">
    {{ block "css" . }}{{ end }}
    </head>
    
<body>
    <header>
        <a href="/">Home</a>
        <a href="/about">About</a>
        <a href="/aboat">A boat</a>
        <a href="/contact">Contact us!</a>
        <hr>
    </header>
    <div>
        {{ block "content" . }}{{ end }}
    </div>
    {{ block "js" . }}{{ end }}
</body>
</html>