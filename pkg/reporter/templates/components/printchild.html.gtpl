
      <!-- print child -->
      {{ define "printchild" }}
      <div class="w-full lg:w-1/2 hover:bg-gray-800 hover:rounded py-1 px-2" onclick="toggleDetails('{{ .Id }}')">
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
      {{ if eq (len .Children) 0}}
      <div class="rounded-md border-gray-800 shadow-inner bg-gray-600 hidden mt-5 mb-5" id="{{ .Id }}">
        <div class="px-4 py-4">
          {{ range .Logs }}
          <p>{{ . }}</p>
          {{ end }}
        </div>
      </div>
      {{ end }}
      <div class="ml-5 lg:ml-10 hidden" id="{{ .Id }}">
        {{ range .Children }}
          {{ template "printchild" . }}
        {{ end }}
      </div>
      {{ end }}