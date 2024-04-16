run:
	@echo "http://localhost:6969"
	@wgo -file=.go -file=.templ -xfile=_templ.go templ generate :: go run main.go
