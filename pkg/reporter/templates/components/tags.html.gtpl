{{ define "tags" }}
<!-- tags -->
<div class="w-11/12 mx-auto flex flex-wrap gap-y-2 text-sm mt-2">
    {{ range .Tags }}
    <div class="border border-gray-700 rounded flex ml-2">
        <div class="border-r border-gray-700 bg-gray-700"><span class="px-2">{{ .Key }}</span></div>
          <div class="border-r border-gray-700"><span class="px-2">{{ .Value }}</span></div>
    </div>
    {{ end }}
</div>
{{ end }}