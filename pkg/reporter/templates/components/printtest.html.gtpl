<!-- print test -->
{{ define "printtest" }}
<div class="mt-0.5 mb-5 border border-gray-700 px-3 py-1 rounded shadow-md hover:shadow-lg hover:border-gray-500 flex " onclick="toggleDetails('{{ .Id }}')">
    <div class="flex">
    <p>
        <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 8.75L14.25 12L10.75 15.25"></path>
        </svg>
    </p>
    <p class="my-auto">{{ .Name }}</p>
    </div>
</div>

<div class="mb-10" id="{{ .Id }}">
<!-- range over tests -->
{{ range .Children }}
{{ template "printchild" . }}
{{ end }}
</div>
{{ end }}