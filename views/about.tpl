{{ template "layout.tpl" . }}
{{ define "content"}}
<head>
    <title>{{ .Title }}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="stylesheet" href="{{.BaseUrl}}/static/css/output.css">
    <link rel="icon" type="image/x-icon" href="/favicon.ico">
    {{ block "css" . }}{{ end }}
    </head>
        <body>
        <h1>About us:</h1>
        <pre>hi
This page has been take by the coporation named jeff and so now its for testing.</pre>
        <button class="btn btn-primary ms-1">Primary</button>
        <button class="btn btn-secondary">Secondary</button>
        <button class="btn btn-accent">Accent</button>
        <button class="btn btn-info">Info</button>
        <button class="btn btn-success">Success</button>
        <button class="btn btn-warning">Warning</button>
        <button class="btn btn-error">Error</button>
        <button class="btn btn-neutral">Neutral</button>

        </body>
{{end}}
