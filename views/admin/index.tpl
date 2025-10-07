{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="{{.BaseUrl}}/static/css/output.css">
    <link rel="icon" type="image/x-icon" href="/favicon.ico">
    {{ block "css" . }}{{ end }}
    </head>
<h1 class="text-4xl">Admin</h1>
<hr>
<br>
<h2 class="text-4xl">Your messages:</h2>
{{ template "admin/profile.tpl" . }}
{{ end }}
