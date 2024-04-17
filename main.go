package main

import (
	"context"
	"log"
	"os"

	"github.com/hyperpuncher/krissmieng/components"
)

func main() {
	f, err := os.Create("krissmieng.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = components.Page().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
