run:
	@echo "http://localhost:6969"
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: tailwindcss-daisyui -i assets/css/input.css -o assets/css/output.css :: go run main.go

build:
	@tailwindcss-daisyui -i assets/css/input.css -o assets/css/output.css --minify
	@go build .
