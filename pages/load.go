package pages

import (
	"text/template"
)

const html = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Crypto Price Tracker</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link rel="preconnect" href="https://fonts.googleapis.com" />
    <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
    <link
      href="https://fonts.googleapis.com/css2?family=IBM+Plex+Mono&family=IBM+Plex+Sans&display=swap"
      rel="stylesheet"
    />
  </head>
  <body>
    <script>
      tailwind.config = {
        theme: {
          fontFamily: {
            mono: ["IBM Plex Mono", "monospace"],
            sans: ["IBM Plex Sans", "sans-serif"],
          },
        },
      };
    </script>
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 m-4 text-black font-sans"
    >
      {{ range $cg := . }}
      <div
        class="flex flex-col gap-2 p-2 bg-neutral-100 justify-center items-center h-48 border border-neutral-200"
      >
        <img
          src="{{ $cg.Image }}"
          alt="crypto"
          class="mx-auto w-12 sm:w-16 aspect-square"
        />
        <div class="flex justify-evenly w-64 sm:w-48 text-lg">
          <p>{{ $cg.Name }}</p>
          <p class="font-bold">{{ $cg.Symbol }}</p>
        </div>
        <div class="flex justify-evenly w-64 sm:w-48 font-mono">
          <p>{{ $cg.CurrentPrice }}$</p>
          <p>{{ $cg.PriceChangePercentage24H }}%</p>
        </div>
      </div>
      {{ end }}
    </div>
  </body>
</html>
`

func LoadTemplate() (*template.Template) {
	tmpl, err := template.New("index").Parse(html)
	if err != nil {
		panic("HTML cannot be parsed")
	}
	return tmpl
}
