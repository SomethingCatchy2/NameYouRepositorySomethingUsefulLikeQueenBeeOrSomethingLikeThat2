{{ template "layout.tpl" . }}
{{ define "content" }}
        <body>
        <h2>{{ .Title }}</h2>
        <p>Ukk Leldredge was here!</p>
            <pre>people who have seen this page:
           {{ .Visits }}
            </pre>
        </body>
{{ end }}