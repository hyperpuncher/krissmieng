package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/hyperpuncher/krissmieng/components"
)

func main() {
	hello := components.Hello()

	http.Handle("/", templ.Handler(hello))

	http.ListenAndServe(":6969", nil)
}
