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
<h1 class="text-4xl">{{.Title}}</h1>
<table class="border px-4 py-2 table w-full mt-4">
  {{if eq (len .List) 0}} 
  <tr>
    <td colspan="3">No records</td> 
  </tr>
  {{else}}

    <tr>
    {{range index .List 0}}
      <th class="bg-gray-200">{{.Key}}</th>
    {{end}}
    </tr>

  <tbody>
    {{range .List}}
    <tr>
    {{range .}}
      {{if eq .Key "Id"}}
        <td class="border px-4 py-2">
        <a class="link link-primary" href="{{$.BaseHref}}/{{.Value}}">
        {{.Value}}
        </a>
        </td>
      </a>
      {{else}}
        <td class="border px-4 py-2">{{.Value}}</td>
      {{end}}
    {{end}}
    </tr>
    </a>
    {{end}}
  </tbody>
  {{end}}
</table>
{{ end }}
