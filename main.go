package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/hyperpuncher/krissmieng/components"
)

func main() {
	assets := http.FileServer(http.Dir("assets"))
	page := components.Page()

	http.Handle("/", templ.Handler(page))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	http.ListenAndServe(":6969", nil)
}
