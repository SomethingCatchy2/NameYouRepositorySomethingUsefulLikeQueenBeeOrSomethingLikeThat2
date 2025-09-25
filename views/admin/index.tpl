{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<h1 class="text-4xl">Admin</h1>
<hr>
<br>
<h2 class="text-4xl">Your messages:</h2>
{{ template "admin/profile.tpl" . }}
{{ end }}
