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
      {{ template "tags" . }}
    </div>

    <!-- templates level report -->
    <div class="mt-5 mb-10">
      <div class="w-11/12 mx-auto lg:ml-16 text-sm cursor-pointer">
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