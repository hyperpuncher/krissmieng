run:
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: tailwindcss-daisyui -i assets/css/input.css -o assets/css/output.css :: go run main.go

build:
	@templ generate
	@tailwindcss-daisyui -i assets/css/input.css -o assets/css/output.css --minify
	@go run main.go
