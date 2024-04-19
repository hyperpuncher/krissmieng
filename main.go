package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/hyperpuncher/krissmieng/components"
)

func main() {
	createHtml()
}

func createHtml() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = components.Page().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func serveHtml() {
	http.Handle("/", templ.Handler(components.Page()))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("serving on http://localhost:6969")
	http.ListenAndServe(":6969", nil)
}
