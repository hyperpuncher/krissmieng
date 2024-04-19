run:
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: unocss :: go run main.go

build:
	@templ generate
	@unocss -m
	@go run main.go
