<!-- checkbox with circle -->
{{ define "check-with-circle" }}
<div class="bg-green-600 rounded-full w-6 h-6 transform scale-75 ml-2">
  <span class="text-white">
    <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
        d="M5.75 12.8665L8.33995 16.4138C9.15171 17.5256 10.8179 17.504 11.6006 16.3715L18.25 6.75" />
    </svg>
  </span>
</div>
{{ end }}

<!-- checkbox with circle -->
{{ define "check" }}
<div class="text-green-400 font-bold w-6 h-6 transform scale-75 ml-2">
  <span>
    <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
        d="M5.75 12.8665L8.33995 16.4138C9.15171 17.5256 10.8179 17.504 11.6006 16.3715L18.25 6.75" />
    </svg>
  </span>
</div>
{{ end }}

<!-- close with circle -->
{{ define "close-with-circle" }}
<div class="bg-red-600 rounded-full w-6 h-6 transform scale-75 ml-2">
  <span class="text-white">
    <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 6.75L6.75 17.25"></path>
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 6.75L17.25 17.25"></path>
    </svg>
  </span>
</div>
{{ end }}

<!-- close -->
{{ define "close" }}
<div class="text-red-400 font-bold w-6 h-6 transform scale-75 ml-2">
  <span>
    <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 6.75L6.75 17.25"></path>
      <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 6.75L17.25 17.25"></path>
    </svg>
  </span>
</div>
{{ end }}

<!-- skip with circle -->
{{ define "skip-with-circle" }}
<div class="bg-gray-600 rounded-full w-6 h-6 transform scale-75 ml-2 flex text-center justify-center">
  <span class="text-white font-bold">
    !
  </span>
</div>
{{ end }}

<!-- skip -->
{{ define "skip" }}
<div class="text-gray-400 font-bold w-6 h-6 transform scale-75 ml-2 text-center justify-center">
  <span>
    !
  </span>
</div>
{{ end }}

<!doctype html>
<html>

  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://cdn.tailwindcss.com"></script>
  </head>

  <body class="bg-gray-900 text-white">
    <!-- navigation -->
    <div class="h-12 bg-gray-800 w-full flex">
      <div class="my-auto ml-2">
        <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
            d="M4.75 5.75H19.25"></path>
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
            d="M4.75 18.25H19.25"></path>
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
            d="M4.75 12H19.25"></path>
        </svg>
      </div>
      <div class="my-auto ml-2 text-sm">
        <span class="bg-gray-700 px-2 py-0.5 rounded text-red-400">{{ .TitlePrimary }}</span>
        {{  .TitleSecondary }}
      </div>
    </div>

    <!-- os selector -->
    <div class="mt-10">
      <div class="w-11/12 mx-auto flex mb-2 border-b border-gray-100 text-sm">
        <div class="ml-5 rounded-t-md bg-gray-700 hover:cursor-pointer px-2 py-2">Linux</div>
      </div>
    </div>

    
    <div class="mt-5">
      
      <!-- tags -->
      <div class="w-11/12 mx-auto flex text-sm">
        {{ range .Tags }}
        <div class="border border-gray-700 rounded flex ml-2">
          <div class="border-r border-gray-700 bg-gray-700"><span class="px-2">{{ .Key }}</span></div>
          <div class="border-r border-gray-700"><span class="px-2">{{ .Value }}</span></div>
        </div>
        {{ end }}
      </div>

      <!-- print grandchild -->
      {{ define "printgrandchild" }}
      <div class="w-1/2 relative flex hover:bg-gray-800 hover:rounded py-1" onclick="toggleLogs('{{ .Id }}')">
          {{ if eq .Result "pass" }}
            {{  template "check" }}
          {{ end }}

          {{ if eq .Result "fail" }}
            {{  template "close" }}
          {{ end }}

          {{ if eq .Result "skip" }}
            {{ template "skip" }}
          {{ end }}

        <div class="w-full ml-1 mr-2 mt-0.5 grid grid-cols-2 gap-4">
          <div class="col-span-1">{{ .Name }}</div>
          <div class="col-span-1 flex justify-end">
            <div class="text-sm">{{ .Duration }}s</div>
            <div class="ml-2 text-green-100 hover:text-green-400">
              <svg width="24" height="24" fill="none" viewBox="0 0 24 24">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 19.25H16.25C17.3546 19.25 18.25 18.3546 18.25 17.25V9L14 4.75H7.75C6.64543 4.75 5.75 5.64543 5.75 6.75V17.25C5.75 18.3546 6.64543 19.25 7.75 19.25Z"></path>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 9.25H13.75V5"></path>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 15.25H14.25"></path>
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 12.25H14.25"></path>
              </svg>
            </div>
          </div>  
        </div>
      </div>
      <!-- logs -->
      <div class="rounded-md border-gray-800 shadow-inner bg-gray-600 hidden mt-5 mb-5 mx-10" id="{{ .Id }}">
        <div class="px-4 py-4">
        {{ range .Logs }}
          <p>{{ . }}</p>
        {{ end }}
        </div>
      </div>
      {{ end }}


      <!-- print child -->
      {{ define "printchild" }}
      <div class="w-1/2 hover:bg-gray-800 hover:rounded py-1 px-2" onclick="toggleDetails('{{ .Id }}')">
        <div class="flex">
          {{ if eq .Result "pass" }}
            {{  template "check-with-circle" }}
          {{ end }}
          
          {{ if eq .Result "fail" }}
            {{  template "close-with-circle" }}
          {{ end }}

          {{ if eq .Result "skip" }}
            {{  template "skip-with-circle" }}
          {{ end }}

          <div class="w-full ml-2 mt-0.5 grid grid-cols-2">
            <div class="col-span-1 text-left"><p>{{ .Name }}</p></div>
            <div><div class="text-sm text-right">{{ .Duration }}s</div></div>
          </div>
        </div>
      </div>
      <!-- logs -->
      {{ if or (eq .Result "skip") (eq .Result "fail") }}
      {{ if eq  (len .Children) 0}}
      <div class="rounded-md border-gray-800 shadow-inner bg-gray-600 hidden mt-5 mb-5" id="{{ .Id }}">
        <div class="px-4 py-4">
          {{ range .Logs }}
          <p>{{ . }}</p>
          {{ end }}
        </div>
      </div>
      {{ end }}
      {{ end }}
      <div class="ml-10 hidden" id="{{ .Id }}">
        {{ range .Children }}
          {{ template "printgrandchild" . }}
        {{ end }}
      </div>
      {{ end }}

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

      <!-- templates level report -->
      <div class="mt-5 mb-10">
        <div class="w-11/12 ml-16 text-sm cursor-pointer">
        {{ range .Tests }}
        <div class="py-1">
          {{ template "printtest" . }}
        </div>
        {{ end }}
        </div>
    </div>
    <script>
      function toggleDetails(id) {
        document.getElementById(id).classList.toggle("hidden");
      }

      function toggleLogs(id) {
        document.getElementById(id).classList.toggle("hidden");
      }
    </script>
  </body>
</html>