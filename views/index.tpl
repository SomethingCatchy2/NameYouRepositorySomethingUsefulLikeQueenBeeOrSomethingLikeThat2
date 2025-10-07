{{ template "layout.tpl" . }}
{{ define "content" }}
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="{{.BaseUrl}}/static/css/output.css">
    <link rel="icon" type="image/x-icon" href="/favicon.ico">
    {{ block "css" . }}{{ end }}
    </head>
<body>
        
        <h2>{{ .Title }}</h2>
        <p>Welcome to the Home Page!</p>
        <p>This page is mostly pointless, please head to a more important page!!</p>
        
        
</body>
{{ end }}