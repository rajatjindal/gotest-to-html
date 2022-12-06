      <!-- print grandchild -->
      {{ define "printgrandchild" }}
      <div class="w-full lg:w-1/2 relative flex hover:bg-gray-800 hover:rounded py-1" onclick="toggleLogs('{{ .Id }}')">
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
      <div class="rounded-md border-gray-800 shadow-inner bg-gray-600 hidden mt-5 mb-5" id="{{ .Id }}">
        <div class="px-4 py-4">
        {{ range .Logs }}
          <p>{{ . }}</p>
        {{ end }}
        </div>
      </div>
      {{ end }}
