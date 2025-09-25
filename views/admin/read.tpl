{{ template "admin/layout.tpl" . }}
{{ define "content" }}
<h1 class="text-4xl m-1">{{.Title}}</h1>
<hr>
<br>
<button class="btn btn-error m-4" onclick="fetch(window.location.href, {method:'DELETE'}).then(()=>{window.location.href = '{{.BaseHref}}'})">Walk the Plank! (Delete)</button>


<table class="border px-4 py-2 table w-full mt-4">

  
    <tr>
      <th class="bg-gray-200">Key</th>
      <th class="bg-gray-200">Value</th>
    </tr>

  <tbody>
    {{range .Item}}

    <tr>
      <td class="border px-4 py-2">{{.Key}}</td>
      <td class="border px-4 py-2">{{.Value}}</td>
    </tr>

    {{end}}
  </tbody>
</table>
{{ end }}
